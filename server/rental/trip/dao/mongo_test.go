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
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			name:       "another_in_progress",
			tripID:     "622c60186800fc9e2ca1480e",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			name:       "in_progress_by_another_account",
			tripID:     "622c60186800fc9e2ca1481e",
			accountID:  "account2",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	for _, cc := range cases {
		mgo.NewObjIDWithValue(id.TripID(cc.tripID))
		tr, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: cc.accountID,
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
		AccountId: acct,
		CarId:     "car1",
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

func TestGetTrips(t *testing.T) {
	rows := []struct {
		id        string
		accountID string
		status    rentalpb.TripStatus
	}{
		{
			id:        "632c60186800fc9e2ca1480d",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "642c60186800fc9e2ca1480d",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "652c60186800fc9e2ca1480d",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "682c60186800fc9e2ca1480d",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "662c60186800fc9e2ca1480d",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			id:        "672c60186800fc9e2ca1480d",
			accountID: "account_id_for_get_trips_1",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
	}
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb : %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	for _, r := range rows {
		mgo.NewObjIDWithValue(id.TripID(r.id))
		_, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: r.accountID,
			Status:    r.status,
		})
		if err != nil {
			t.Fatalf("cannot create rows : %v", err)
		}

	}

	cases := []struct {
		name       string
		accountID  string
		status     rentalpb.TripStatus
		wantCount  int
		wantOnlyID string
	}{
		{
			name:      "get_all",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCount: 5,
		},
		{
			name:      "get_in_progress",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_IN_PROGRESS,
			wantCount: 1,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			tr, err := m.GetTrips(context.Background(), id.AccountID(cc.accountID), cc.status)
			if err != nil {
				t.Errorf("cannot get trips: %v", err)
			}
			if cc.wantCount != len(tr) {
				t.Errorf("incorrect result count:want:%d,got:%d", cc.wantCount, len(tr))
			}
			if cc.wantOnlyID != "" && len(tr) > 0 {
				if cc.wantOnlyID != tr[0].ID.Hex() {
					t.Errorf("only_id incorrect;want:%q,got:%q", cc.wantOnlyID, tr[0].ID.Hex())
				}
			}
		})
	}
}

func TestUpdateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	tid := id.TripID("662c60186800fc9e2ca1480d")
	var now int64 = 10000
	mgo.NewObjIDWithValue(tid)
	mgo.UpdateAt = func() int64 {
		return now
	}
	aid := id.AccountID("account_for_update")
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi",
		},
	})
	if err != nil {
		t.Fatalf("cannot create trip:%v", err)
	}
	if tr.UpdatedAt != 10000 {
		t.Fatalf("wrong updateat;want:10000,got:%d", tr.UpdatedAt)
	}
	update := &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi_updated",
		},
	}
	cases := []struct {
		name           string
		now            int64
		withUppdatedAt int64
		wantErr        bool
	}{
		{
			name:           "normal_update",
			now:            20000,
			withUppdatedAt: 10000,
		},
		{
			name:           "update_with_stale_timestamp",
			now:            30000,
			withUppdatedAt: 10000,
			wantErr:        true,
		},
		{
			name:           "update_with_refetch",
			now:            40000,
			withUppdatedAt: 20000,
		},
	}

	for _, cc := range cases {
		// 用case的now替代mgo.updatedAt
		now = cc.now
		err := m.UpdateTrip(c, tid, aid, cc.withUppdatedAt, update)
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s:want error: got none", cc.name)
			} else {
				continue
			}
		}
		updatedTrip, err := m.GetTrip(c, tid, aid)
		if err != nil {
			t.Errorf("%s:cannot get trip after update:%v", cc.name, err)
		}
		if cc.now != updatedTrip.UpdatedAt {
			t.Errorf("%s:incorrect updatedat:want %d,got %d", cc.name, cc.now, updatedTrip.UpdatedAt)
		}

	}
}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
