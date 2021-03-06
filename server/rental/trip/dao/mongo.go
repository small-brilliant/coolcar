package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField      = "trip"
	accountIdField = tripField + ".accountid"
	statusField    = tripField + ".status"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}

type TripRecord struct {
	mgo.IDField      `bson:"inline"`
	mgo.UpdatedField `bson:"inline"`
	Trip             *rentalpb.Trip `bson:"trip"`
}

// TODO: 同一个account最多只能有一个进行中的trip
// TODO: 强类型化tripID
// TODO: 表格驱动测试

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	// 这两个参数是不固定的，交给mgo管理，在测试的时候可以固定
	r.ID = mgo.NewObjID()
	r.UpdatedAt = mgo.UpdateAt()
	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (m *Mongo) GetTrip(c context.Context, id id.TripID, accountID id.AccountID) (*TripRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}
	res := m.col.FindOne(c, bson.M{
		mgo.IDFieldName: objID,
		accountIdField:  accountID,
	})
	if err := res.Err(); err != nil {
		return nil, err
	}
	var tr TripRecord
	err = res.Decode(&tr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode : %v", err)
	}
	return &tr, err
}

// GetTrips gets trips for the account by status.
// If status is not specified, gets all trips for the account.
func (m *Mongo) GetTrips(c context.Context, accountID id.AccountID, status rentalpb.TripStatus) ([]*TripRecord, error) {
	filter := bson.M{
		accountIdField: accountID.String(),
	}
	if status != rentalpb.TripStatus_TS_NOT_SPECIFIED {
		filter[statusField] = status
	}

	res, err := m.col.Find(c, filter)
	if err != nil {
		return nil, err
	}
	var trips []*TripRecord
	for res.Next(c) {
		var trip TripRecord
		err := res.Decode(&trip)
		// 其中一行出错了的处理方式
		if err != nil {
			return nil, err
		}
		trips = append(trips, &trip)
	}
	return trips, nil
}

func (m *Mongo) UpdateTrip(c context.Context, tripid id.TripID, accountID id.AccountID, updatedAt int64, trip *rentalpb.Trip) error {
	objID, err := objid.FromID(tripid)
	if err != nil {
		return fmt.Errorf("invalid id : %v", err)
	}
	newUpdatedAt := mgo.UpdateAt()
	res, err := m.col.UpdateOne(c, bson.M{
		mgo.IDFieldName:        objID,
		accountIdField:         accountID.String(),
		mgo.UpdatedAtFieldName: updatedAt,
	}, mgo.Set(bson.M{
		tripField:              trip,
		mgo.UpdatedAtFieldName: newUpdatedAt,
	}))
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
