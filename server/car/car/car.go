package car

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/dao"
	"coolcar/shared/id"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Mongo     *dao.Mongo
	Logger    *zap.Logger
	Publisher Publisher
}
type Publisher interface {
	Publish(context.Context, *carpb.CarEntity) error
}

func (s *Service) CreateCar(c context.Context, req *carpb.CreateCarRequest) (*carpb.CarEntity, error) {
	cr, err := s.Mongo.CreateCar(c)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &carpb.CarEntity{
		Id:  cr.ID.Hex(),
		Car: cr.Car,
	}, nil
}
func (s *Service) GetCar(c context.Context, req *carpb.GetCarRequest) (*carpb.Car, error) {
	cr, err := s.Mongo.GetCar(c, id.CarID(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return cr.Car, nil
}
func (s *Service) GetCars(c context.Context, req *carpb.GetCarsRequest) (*carpb.GetCarsResponse, error) {
	cars, err := s.Mongo.GetCars(c)
	if err != nil {
		s.Logger.Error("cannot found cars", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	res := &carpb.GetCarsResponse{}
	for _, cr := range cars {
		res.Cars = append(res.Cars, &carpb.CarEntity{
			Id:  cr.ID.Hex(),
			Car: cr.Car,
		})
	}
	return res, nil
}
func (s *Service) LockCar(c context.Context, req *carpb.LockCarRequest) (*carpb.LockCarResponse, error) {
	car, err := s.Mongo.UpdateCar(c, id.CarID(req.Id), carpb.CarStatus_UNLOCKED, &dao.CarUpdate{
		Status: carpb.CarStatus_LOCKING,
	})
	if err != nil {
		code := codes.Internal
		if err == mongo.ErrNoDocuments {
			code = codes.NotFound
		}
		return nil, status.Errorf(code, "cannot update: %v", err)
	}
	s.publish(c, car)
	return &carpb.LockCarResponse{}, nil
}
func (s *Service) UnlockCar(c context.Context, req *carpb.UnlockCarRequest) (*carpb.UnlockCarResponse, error) {
	car, err := s.Mongo.UpdateCar(c, id.CarID(req.Id), carpb.CarStatus_LOCKED, &dao.CarUpdate{
		Status:       carpb.CarStatus_UNLOCKING,
		Driver:       req.Driver,
		UpdateTripID: true,
		TripID:       id.TripID(req.TripId),
	})
	if err != nil {
		code := codes.Internal
		if err == mongo.ErrNoDocuments {
			code = codes.NotFound
		}
		return nil, status.Errorf(code, "cannot update: %v", err)
	}
	s.publish(c, car)
	return &carpb.UnlockCarResponse{}, nil
}

// 是车上的软件来调用的
func (s *Service) UpdateCar(c context.Context, req *carpb.UpdateCarRequest) (*carpb.UpdateCarResponse, error) {
	update := &dao.CarUpdate{
		Status:   req.Status,
		Position: req.Position,
	}

	if req.Status == carpb.CarStatus_LOCKED {
		update.Driver = &carpb.Driver{}
		update.TripID = id.TripID("")
		update.UpdateTripID = true
	}
	car, err := s.Mongo.UpdateCar(c, id.CarID(req.Id), carpb.CarStatus_CS_NOT_SPECIFIED, update)
	if err != nil {
		code := codes.Internal
		if err == mongo.ErrNoDocuments {
			code = codes.NotFound
		}
		return nil, status.Errorf(code, "cannot update: %v", err)
	}
	s.publish(c, car)
	return &carpb.UpdateCarResponse{}, nil
}

func (s *Service) publish(c context.Context, car *dao.CarRecord) {
	err := s.Publisher.Publish(c, &carpb.CarEntity{
		Id:  car.ID.Hex(),
		Car: car.Car,
	})
	if err != nil {
		s.Logger.Warn("cannot publish", zap.Error(err))
	}
}