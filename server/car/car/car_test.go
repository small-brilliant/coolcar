package car

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/dao"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	mongotesting "coolcar/shared/mongo/testing"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestCarUpdate(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger : %v", err)
	}
	s := &Service{
		Logger:    logger,
		Mongo:     dao.NewMongo(mc.Database("coolcar")),
		Publisher: &testPublisher{},
	}
	carID := id.CarID("521c60186800fc9e2ca14801")
	mgo.NewObjIDWithValue(carID)
	car1, err := s.CreateCar(c, &carpb.CreateCarRequest{})
	if err != nil {
		t.Fatalf("cannot create car: %v", err)
	}
	fmt.Println(car1)
	cases := []struct {
		name    string
		op      func() error
		want    string
		wantErr bool
	}{
		{
			name: "get_car",
			op: func() error {
				return nil
			},
			want: `{"status":1,"position":{"latitude":30,"longitude":120}}`,
		},
		{
			name: "unlock_car",
			op: func() error {
				_, err := s.UnlockCar(c, &carpb.UnlockCarRequest{
					Id:     carID.String(),
					TripId: "test_tripID",
					Driver: &carpb.Driver{
						Id:        "test_driver",
						AvatarUrl: "test_avatarUrl",
					},
				})
				return err
			},
			want: `{"status":2,"driver":{"id":"test_driver","avatar_url":"test_avatarUrl"},"position":{"latitude":30,"longitude":120},"trip_id":"test_tripID"}`,
		},
		{
			name: "unlock_complete",
			op: func() error {
				_, err := s.UpdateCar(c, &carpb.UpdateCarRequest{
					Id: carID.String(),
					Position: &carpb.Location{
						Latitude:  31,
						Longitude: 120,
					},
					Status: carpb.CarStatus_UNLOCKED,
				})
				return err
			},
			want: `{"status":3,"driver":{"id":"test_driver","avatar_url":"test_avatarUrl"},"position":{"latitude":31,"longitude":120},"trip_id":"test_tripID"}`,
		},
		{
			name: "unlock_car_by_another_driver",
			op: func() error {
				_, err := s.UnlockCar(c, &carpb.UnlockCarRequest{
					Id:     carID.String(),
					TripId: "bad_tripID",
					Driver: &carpb.Driver{
						Id:        "bad_driver",
						AvatarUrl: "bad_avatarUrl",
					},
				})
				return err
			},
			wantErr: true,
		},
		{
			name: "lock_car",
			op: func() error {
				_, err := s.LockCar(c, &carpb.LockCarRequest{
					Id: carID.String(),
				})
				return err
			},
			want: `{"status":4,"driver":{"id":"test_driver","avatar_url":"test_avatarUrl"},"position":{"latitude":31,"longitude":120},"trip_id":"test_tripID"}`,
		},
		{
			name: "lock_complete",
			op: func() error {
				_, err := s.UpdateCar(c, &carpb.UpdateCarRequest{
					Id:     carID.String(),
					Status: carpb.CarStatus_LOCKED,
				})
				return err
			},
			want: `{"status":1,"driver":{},"position":{"latitude":31,"longitude":120}}`,
		},
	}

	for _, cc := range cases {
		err := cc.op()
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want error; got node", cc.name)
			} else {
				continue
			}
		}
		car, err := s.GetCar(c, &carpb.GetCarRequest{
			Id: carID.String(),
		})
		if err != nil {
			t.Errorf("%s: operation faild: %v", cc.name, err)
			continue
		}
		b, err := json.Marshal(car)
		if err != nil {
			t.Errorf("%s:faild marshal response: %v", cc.name, err)
		}
		got := string(b)
		if cc.want != got {
			t.Errorf("%s:incorrect response; want: %s, got: %s", cc.name, cc.want, got)
		}
	}
}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}

type testPublisher struct{}

func (p *testPublisher) Publish(c context.Context, car *carpb.CarEntity) error {
	return nil
}
