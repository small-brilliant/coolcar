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
	"math/rand"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	pm := &profileManager{}
	cm := &carManager{}
	s := newService(c, t, pm, cm)
	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Location{
			Latitude:  32.12,
			Longitude: 114.2555,
		},
	}
	pm.iID = "identity1"
	golden := `{"account_id":%q,"car_id":"car1","start":{"location":{"latitude":32.12,"longitude":114.2555},"poi_name":"天安门","timestamp_sec":1647432905},"current":{"location":{"latitude":32.12,"longitude":114.2555},"poi_name":"天安门","timestamp_sec":1647432905},"status":1,"identity_id":"identity1"}`
	nowFunc = func() int64 {
		return 1647432905
	}
	cases := []struct {
		name         string
		accountID    string
		tripID       string
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:      "normal_create",
			accountID: "account1",
			tripID:    "662c60186800fc9e2ca1480d",
			want:      fmt.Sprintf(golden, "account1"),
		},
		{
			name:       "profile_err",
			accountID:  "account2",
			tripID:     "662c60186800fc9e2ca14801",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			accountID:    "account3",
			tripID:       "662c60186800fc9e2ca14802",
			carVerifyErr: fmt.Errorf("verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err",
			accountID:    "account4",
			tripID:       "662c60186800fc9e2ca14803",
			carUnlockErr: fmt.Errorf("unlock"),
			wantErr:      false, // 解锁失败，创建trip还是成功的
			want:         fmt.Sprintf(golden, "account4"),
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			c = auth.ContestWithAccontId(c, id.AccountID(cc.accountID))
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

func (cm *carManager) Verfigy(context.Context, id.CarID, *rentalpb.Location) error {
	return cm.verifyErr
}
func (cm *carManager) Unlock(c context.Context, cid id.CarID, aid id.AccountID, tid id.TripID, avatarURL string) error {
	return cm.unlockErr
}
func (cm *carManager) Lock(c context.Context, cid id.CarID) error {
	return cm.unlockErr
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

// 测试trip的整个流程
func TestTripLifecycle(t *testing.T) {
	c := auth.ContestWithAccontId(context.Background(), id.AccountID("account_id_lifecycle"))
	s := newService(c, t, &profileManager{}, &carManager{})
	tid := id.TripID("621c60186800fc9e2ca14801")
	mgo.NewObjIDWithValue(tid)
	cases := []struct {
		name string
		now  int64
		op   func() (*rentalpb.Trip, error)
		want string
	}{
		{
			name: "create_trip",
			now:  10000,
			op: func() (*rentalpb.Trip, error) {
				e, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{
					Start: &rentalpb.Location{
						Latitude:  34,
						Longitude: 122,
					},
					CarId: "car123",
				})
				if err != nil {
					return nil, err
				}
				return e.Trip, nil
			},
			want: `{"account_id":"account_id_lifecycle","car_id":"car123","start":{"location":{"latitude":34,"longitude":122},"poi_name":"中关村","timestamp_sec":10000},"current":{"location":{"latitude":34,"longitude":122},"poi_name":"中关村","timestamp_sec":10000},"status":1}`,
		},
		{
			name: "update_trip",
			now:  20000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
					Current: &rentalpb.Location{
						Latitude:  28.2,
						Longitude: 122,
					},
				})
			},
			want: `{"account_id":"account_id_lifecycle","car_id":"car123","start":{"location":{"latitude":34,"longitude":122},"poi_name":"中关村","timestamp_sec":10000},"current":{"location":{"latitude":28.2,"longitude":122},"fee_cent":9116,"km_driven":359.56460921309167,"poi_name":"天安门","timestamp_sec":20000},"status":1}`,
		},
		{
			name: "finish_trip",
			now:  30000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id:      tid.String(),
					EndTrip: true,
				})
			},
			want: `{"account_id":"account_id_lifecycle","car_id":"car123","start":{"location":{"latitude":34,"longitude":122},"poi_name":"中关村","timestamp_sec":10000},"current":{"location":{"latitude":28.2,"longitude":122},"fee_cent":17638,"km_driven":480.5067823692563,"poi_name":"天安门","timestamp_sec":30000},"end":{"location":{"latitude":28.2,"longitude":122},"fee_cent":17638,"km_driven":480.5067823692563,"poi_name":"天安门","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "query_trip",
			now:  40000,
			op: func() (*rentalpb.Trip, error) {
				return s.GetTrip(c, &rentalpb.GetTripRequest{
					Id: tid.String(),
				})
			},
			want: `{"account_id":"account_id_lifecycle","car_id":"car123","start":{"location":{"latitude":34,"longitude":122},"poi_name":"中关村","timestamp_sec":10000},"current":{"location":{"latitude":28.2,"longitude":122},"fee_cent":17638,"km_driven":480.5067823692563,"poi_name":"天安门","timestamp_sec":30000},"end":{"location":{"latitude":28.2,"longitude":122},"fee_cent":17638,"km_driven":480.5067823692563,"poi_name":"天安门","timestamp_sec":30000},"status":2}`,
		},
	}
	rand.Seed(1234)
	for _, cc := range cases {
		nowFunc = func() int64 {
			return cc.now
		}
		trip, err := cc.op()
		if err != nil {
			t.Errorf("%s: operation failed:%v", cc.name, err)
			continue
		}
		b, err := json.Marshal(trip)
		if err != nil {
			t.Errorf("%s:failed marshalling response: %v", cc.name, err)
		}
		got := string(b)
		if cc.want != got {
			t.Errorf("incorrect response:want %s , got %s", cc.want, got)
		}
	}
}
func newService(c context.Context, t *testing.T, pm ProfileManager, cm CarManager) *Service {
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger : %v", err)
	}
	db := mc.Database("coolcar")
	mongotesting.SetupIndex(c, db)
	return &Service{
		ProfileManager: pm,
		CarManager:     cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(db),
		Logger:         logger,
	}
}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
