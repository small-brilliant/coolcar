package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	"coolcar/shared/mongo/objid"
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	CarManager     CarManager
	ProfileManager ProfileManager
	POIManager     POIManager
	Mongo          *dao.Mongo
	Logger         *zap.Logger
}

// ProfileManager defines the ACL(Anti Corruptino Layer)
// for profile verification logic.
type ProfileManager interface {
	Verify(context.Context, id.AccountID) (id.IdentiTyID, error)
}

type CarManager interface {
	// 加入人的位置
	// Verfigy来判断车可不可以被开锁
	Verfigy(c context.Context, cid id.CarID, loc *rentalpb.Location) error
	Unlock(c context.Context, cid id.CarID, aid id.AccountID, tid id.TripID, avatarURL string) error
	Lock(c context.Context, cid id.CarID) error
}

// resolves POI(Point Of Interest)
type POIManager interface {
	Resolve(context.Context, *rentalpb.Location) (string, error)
}

//  CreateTrip(context.Context, *CreateTripRequest) (*TripEntity, error)
// 	GetTrip(context.Context, *GetTripRequest) (*Trip, error)
// 	GetTrips(context.Context, *GetTripsRequest) (*GetTripsResponse, error)
// 	UpdateTrip(context.Context, *UpdateTripRequest) (*Trip, error)
func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	if req.CarId == "" || req.Start == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}
	// 验证驾驶证身份
	iID, err := s.ProfileManager.Verify(c, aid)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	// 检查车辆状态，车子可能已经被租了
	carID := id.CarID(req.CarId)
	err = s.CarManager.Verfigy(c, carID, req.Start)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	// 创建行程：写入数据库，开始计费
	ls := s.calcCurrentStatus(c, &rentalpb.LocationStatus{
		Location:     req.Start,
		TimestampSec: nowFunc(),
	}, req.Start)

	tr, err := s.Mongo.CreateTrip(c, &rentalpb.Trip{
		AccountId:  aid.String(),
		CarId:      carID.String(),
		IdentityId: iID.String(),
		Status:     rentalpb.TripStatus_IN_PROGRESS,
		Start:      ls,
		Current:    ls,
	})

	if err != nil {
		s.Logger.Warn("cannot create trip", zap.Error(err))
		return nil, status.Error(codes.AlreadyExists, "")
	}

	// 车辆开锁
	// 后台开锁，因为不管开锁是否成功，行程都已经创建了
	go func() {
		err = s.CarManager.Unlock(context.Background(), carID, aid, objid.ToTripID(tr.ID), req.AvatarUrl)
		if err != nil {
			s.Logger.Error("cannot unlock car", zap.Error(err))
		}
	}()
	return &rentalpb.TripEntity{
		Id:   tr.ID.Hex(),
		Trip: tr.Trip,
	}, nil
}

func (s *Service) GetTrip(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	fmt.Print("start GetTrip")
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	tr, err := s.Mongo.GetTrip(c, id.TripID(req.Id), aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return tr.Trip, nil
}

func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	trips, err := s.Mongo.GetTrips(c, aid, req.Status)
	if err != nil {
		s.Logger.Error("cannot get trips", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	res := &rentalpb.GetTripsResponse{}
	for _, tr := range trips {
		res.Trips = append(res.Trips, &rentalpb.TripEntity{
			Id:   tr.ID.Hex(),
			Trip: tr.Trip,
		})
	}
	return res, nil
}

func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	tr, err := s.Mongo.GetTrip(c, id.TripID(req.Id), aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	if tr.Trip.Current == nil {
		s.Logger.Error("trip without current set", zap.String("id", tr.ID.String()))
		return nil, status.Error(codes.Internal, "")
	}
	// 考虑到车子不动的情况
	cur := tr.Trip.Current.Location
	if req.Current != nil {
		cur = req.Current
	}
	ls := s.calcCurrentStatus(c, tr.Trip.Current, cur)
	tr.Trip.Current = ls
	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
		// 先关锁在结束订单.
		err := s.CarManager.Lock(c, id.CarID(tr.Trip.CarId))
		if err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "cannot lock car: %v", err)
		}
	}
	err = s.Mongo.UpdateTrip(c, id.TripID(req.Id), aid, tr.UpdatedAt, tr.Trip)
	if err != nil {
		return nil, status.Error(codes.Aborted, "")
	}
	return tr.Trip, nil
}

var nowFunc = func() int64 {
	return time.Now().Unix()
}

const centsPerSec = 0.7
const kmPerSec = 0.02

func (s *Service) calcCurrentStatus(c context.Context, last *rentalpb.LocationStatus, cur *rentalpb.Location) *rentalpb.LocationStatus {
	now := nowFunc()
	elapsedSec := float64(now - last.TimestampSec)

	poi, err := s.POIManager.Resolve(c, cur)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location", cur), zap.Error(err))
	}
	return &rentalpb.LocationStatus{
		Location:     cur,
		FeeCent:      last.FeeCent + int32(centsPerSec*elapsedSec*2*rand.Float64()),
		KmDriven:     last.KmDriven + kmPerSec*elapsedSec*2*rand.Float64(),
		TimestampSec: now,
		PoiName:      poi,
	}
}

func (s *Service) finshTrip(c context.Context, req *rentalpb.UpdateTripRequest) error {
	return nil
}
