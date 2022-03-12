package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connection db: %v", err)
	}
	db := mc.Database("coolcar")
	err = mongotesting.SetupIndex(c, db)
	if err != nil {
		t.Fatalf("cannot SetupIndex : %v", err)
	}
	m := NewMongo(db)

	cases := []struct {
		name       string
		tripID     string
		accountID  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finised",
			tripID:     "622c60186800fc9e2ca1480f",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "another_finised",
			tripID:     "622c60186800fc9e2ca1480d",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "in_progress",
			tripID:     "622c60186800fc9e2ca1480c",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRES,
		},
		{
			name:       "another_in_progress",
			tripID:     "622c60186800fc9e2ca1480e",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRES,
			wantErr:    true,
		},
		{
			name:       "in_progress_by_another_account",
			tripID:     "622c60186800fc9e2ca1481e",
			accountID:  "account2",
			tripStatus: rentalpb.TripStatus_IN_PROGRES,
		},
	}

	for _, cc := range cases {
		mgo.NewObjID = func() primitive.ObjectID {
			return objid.MustFromID(id.TripID(cc.tripID))
		}
		tr, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountID: cc.accountID,
			Status:    cc.tripStatus,
		})
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s:error expected; got none", cc.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s:error creating trip: %v", cc.name, err)
			continue
		}

		if tr.ID.Hex() != cc.tripID {
			t.Errorf("%s: incorrect trip id;want: %q;got:%q", cc.name, cc.tripID, tr.ID.Hex())
		}
	}

}
func TestGetTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewDefalutClient(c)
	if err != nil {
		t.Fatalf("cannot connection db: %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	acct := "account1"
	// 如果运行整个package的时候，上面的TestCreateTrip，以及固定住了ObjID,，所以要在这里进行重置
	mgo.NewObjID = primitive.NewObjectID
	trip, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountID: acct,
		CarID:     "car1",
		Start: &rentalpb.LocationStatus{
			PoiName: "startpoint",
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
		},
		End: &rentalpb.LocationStatus{
			PoiName:  "endpoint",
			FeeCent:  10000,
			KmDriven: 100000,
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 125,
			},
		},
		Current: &rentalpb.LocationStatus{
			PoiName:  "currentpoint",
			FeeCent:  10000,
			KmDriven: 100000,
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 125,
			},
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Errorf("cannot create trip: %v", err)
	}
	// +v 用key：value显示
	// t.Errorf("insert row %s with updateat %v", trip.ID, trip.UpdatedAt)
	got, err := m.GetTrip(c, objid.ToTripID(trip.ID), id.AccountID(acct))
	// got.Trip.Current.PoiName = "badPoiName" 测试 cmp.Diff
	if err != nil {
		t.Errorf("cannot get trip : %v", err)
	}
	if diff := cmp.Diff(trip, got, protocmp.Transform()); diff != "" {
		t.Errorf("result differs;-want +got: %s", diff)
	}
}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
