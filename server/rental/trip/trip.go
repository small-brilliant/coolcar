package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
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

	s.Logger.Info("creat trip", zap.String("start", req.Start), zap.String("accoundID:", aid.String()))
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Service) GetTrip(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	s.Logger.Info("GetTrip", zap.String("accoundID:", aid.String()))
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	s.Logger.Info("GetTrips", zap.String("accoundID:", aid.String()))
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	s.Logger.Info("UpdateTrip", zap.String("accoundID:", aid.String()))
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
