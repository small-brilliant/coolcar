package ai

import (
	"context"
	aipb "coolcar/ai/api/gen/v1"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/mq"
	"math"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	Logger    *zap.Logger
	Publisher mq.Publisher
}

var carMap = make(map[string]bool, 0)

func (s *Service) LicIdentity(c context.Context, req *aipb.LicIdentityRequest) (*aipb.Identity, error) {
	return &aipb.Identity{
		LicNumber:       "12389721398791",
		Name:            "谭谭",
		Gender:          aipb.Gender_MALE,
		BirthDateMillis: 23498237948237,
	}, nil
}
func (s *Service) MeasureDistance(c context.Context, req *aipb.MeasureDistanceRequest) (*aipb.MeasureDistanceResponse, error) {
	start, end := req.From, req.To
	dis := math.Sqrt(math.Pow((end.Latitude-start.Latitude), 2)+math.Pow((end.Longitude-end.Longitude), 2)) * 111
	return &aipb.MeasureDistanceResponse{
		DistanceKm: float32(dis),
	}, nil
}
func (s *Service) SimulateCarPos(c context.Context, req *aipb.SimulateCarPosRequest) (*aipb.SimulateCarPosResponse, error) {
	carMap[req.CarId] = true
	go func(cm map[string]bool) {
		for cm[req.CarId] {
			req.Pos.Latitude += 0.0001
			req.Pos.Longitude += 0.00005
			err := s.Publisher.Publish(c, &carpb.CarEntity{
				Id: req.CarId,
				Car: &carpb.Car{
					Position: &carpb.Location{
						Latitude:  req.Pos.Latitude,
						Longitude: req.Pos.Longitude,
					},
				},
			})
			if err != nil {
				s.Logger.Warn("cannot publish", zap.Error(err))
			}
			time.Sleep(1 * time.Second)
		}
	}(carMap)
	return &aipb.SimulateCarPosResponse{}, nil
}
func (s *Service) EndSimulateCarPos(c context.Context, req *aipb.EndSimulateCarPosRequest) (*aipb.EndSimulateCarPosResponse, error) {
	carMap[req.CarId] = false
	s.Publisher.Publish(c, &carpb.CarEntity{
		Id: req.CarId,
		Car: &carpb.Car{
			Status: carpb.CarStatus_LOCKED,
		},
	})
	return &aipb.EndSimulateCarPosResponse{}, nil
}
