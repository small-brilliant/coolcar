package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	"fmt"

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
	Verfigy(context.Context, id.CarID, *rentalpb.Location) error
	Unlock(context.Context, id.CarID) error
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
	poi, err := s.POIManager.Resolve(c, req.Start)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location", req.Start), zap.Error(err))
	}
	ls := &rentalpb.LocationStatus{
		Location: req.Start,
		PoiName:  poi,
	}
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
		err = s.CarManager.Unlock(context.Background(), carID)
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
	if req.Current != nil {
		tr.Trip.Current = s.calcCurrentStatus(tr.Trip, req.Current)
	}
	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
	}
	s.Mongo.UpdateTrip(c, id.TripID(req.Id), aid, tr.UpdatedAt, tr.Trip)
	s.Logger.Info("UpdateTrip", zap.String("accoundID:", aid.String()))
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Service) calcCurrentStatus(trip *rentalpb.Trip, cur *rentalpb.Location) *rentalpb.LocationStatus {
	return nil
}
