package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/client/poi"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	mongotesting "coolcar/shared/mongo/testing"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestCreateTrip(t *testing.T) {
	c := auth.ContestWithAccontId(context.Background(), id.AccountID("accountID1"))
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger : %v", err)
	}
	pm := &profileManager{}
	cm := &carManager{}
	s := &Service{
		ProfileManager: pm,
		CarManager:     cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		Logger:         logger,
	}
	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Location{
			Latitude:  32.12,
			Longitude: 114.2555,
		},
	}
	pm.iID = "identity1"
	golden := `{"accountID":"accountID1","carID":"car1","start":{"location":{"latitude":32.12,"longitude":114.2555},"poi_name":"天安门"},"current":{"location":{"latitude":32.12,"longitude":114.2555},"poi_name":"天安门"},"status":1,"identity_id":"identity1"}`
	cases := []struct {
		name         string
		tripID       string
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:   "normal_create",
			tripID: "662c60186800fc9e2ca1480d",
			want:   golden,
		},
		{
			name:       "profile_err",
			tripID:     "662c60186800fc9e2ca14801",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			tripID:       "662c60186800fc9e2ca14802",
			carVerifyErr: fmt.Errorf("verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err",
			tripID:       "662c60186800fc9e2ca14803",
			carUnlockErr: fmt.Errorf("unlock"),
			wantErr:      false, // 解锁失败，创建trip还是成功的
			want:         golden,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			mgo.NewObjIDWithValue(id.TripID(cc.tripID))
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyErr
			res, err := s.CreateTrip(c, req)
			if cc.wantErr {
				if err == nil {
					t.Errorf("want error got none")
				} else {
					return
				}
			}

			if err != nil {
				t.Errorf("error creating trip: %v", err)
				return
			}
			if res.Id != cc.tripID {
				t.Errorf("incorrect id; want %q,got %q", cc.tripID, res.Id)
			}
			b, err := json.Marshal(res.Trip)
			if err != nil {
				t.Errorf("cannot marshal response:%v", err)
			}
			tripnStr := string(b)
			if cc.want != tripnStr {
				t.Errorf("incorrect response:want %s , got %s", cc.want, tripnStr)
			}

		})
	}
}

type profileManager struct {
	iID id.IdentiTyID
	err error
}

func (p *profileManager) Verify(context.Context, id.AccountID) (id.IdentiTyID, error) {
	return p.iID, p.err
}

type carManager struct {
	verifyErr error
	unlockErr error
}

func (c *carManager) Verfigy(context.Context, id.CarID, *rentalpb.Location) error {
	return c.verifyErr
}
func (c *carManager) Unlock(context.Context, id.CarID) error {
	return c.unlockErr
}
func TestGetTripWithTripID(t *testing.T) {
	c := auth.ContestWithAccontId(context.Background(), id.AccountID("6231952cb3b93c934db10058"))
	mc, err := mongotesting.NewDefalutClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger : %v", err)
	}
	pm := &profileManager{}
	cm := &carManager{}
	s := &Service{
		ProfileManager: pm,
		CarManager:     cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		Logger:         logger,
	}
	t.Run("TestGetTripWithTripID", func(t *testing.T) {
		req := &rentalpb.GetTripRequest{
			Id: "6231952f73dc3b1260e5610f",
		}
		_, err := s.GetTrip(c, req)
		if err != nil {
			t.Errorf("error geting trip: %v", err)
			return
		}
	})
}
func TestGetTrip(t *testing.T) {
	c := auth.ContestWithAccontId(context.Background(), id.AccountID("accountID123"))
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger : %v", err)
	}
	pm := &profileManager{}
	cm := &carManager{}
	s := &Service{
		ProfileManager: pm,
		CarManager:     cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		Logger:         logger,
	}
	mgo.NewObjIDWithValue(id.TripID("662c60186800fc9e2ca14801"))
	_, err = s.CreateTrip(c, &rentalpb.CreateTripRequest{
		Start: &rentalpb.Location{
			Latitude:  34,
			Longitude: 122,
		},
		CarId: "car123",
	})
	if err != nil {
		t.Errorf("error creating trip: %v", err)
		return
	}
	cases := []struct {
		name      string
		tripID    string
		accountID string
		wantErr   bool
	}{
		{
			name:      "normal_get_trip",
			tripID:    "662c60186800fc9e2ca14801",
			accountID: "account123",
		},
		{
			name:      "bad_get_trip",
			tripID:    "662c60186800fc9e2ca14802",
			accountID: "account123",
			wantErr:   true,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			req := &rentalpb.GetTripRequest{
				Id: cc.tripID,
			}
			_, err := s.GetTrip(c, req)
			if cc.wantErr {
				if err == nil {
					t.Errorf("want error got none")
				} else {
					return
				}
			}
			if err != nil {
				t.Errorf("error geting trip: %v", err)
				return
			}
		})
	}

}
func TestGetTrips(t *testing.T) {

}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
