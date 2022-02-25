package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

type Service struct{}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "dfe",
			DurationSec: 3600,
			FeeCent:     10000,
		},
	}, nil
}
