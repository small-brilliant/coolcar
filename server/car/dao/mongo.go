package dao

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	carField      = "car"
	statusField   = carField + ".status"
	driverField   = carField + ".driver"
	positionField = carField + ".position"
	tripIdField   = carField + ".tripid"
)

type Mongo struct {
	col *mongo.Collection
}
type CarUpdate struct {
	Status       carpb.CarStatus
	Position     *carpb.Location
	Driver       *carpb.Driver
	UpdateTripID bool
	TripID       id.TripID
}
type CarRecord struct {
	mgo.IDField `bson:"inline"`
	Car         *carpb.Car `bson:"car"`
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("car"),
	}
}
func (m *Mongo) CreateCar(c context.Context) (*CarRecord, error) {
	r := &CarRecord{
		Car: &carpb.Car{
			Position: &carpb.Location{
				Latitude:  32.64709,
				Longitude: 117.01743,
			},
			Status: carpb.CarStatus_LOCKED,
		},
	}
	r.ID = mgo.NewObjID()
	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (m *Mongo) GetCar(c context.Context, carId id.CarID) (*CarRecord, error) {
	objId, err := objid.FromID(carId)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}
	res := m.col.FindOne(c, bson.M{
		mgo.IDFieldName: objId,
	})
	return convertSingleResult(res)
}
func (m *Mongo) GetCars(c context.Context) ([]*CarRecord, error) {
	filter := bson.M{}
	res, err := m.col.Find(c, filter, options.Find())
	if err != nil {
		return nil, err
	}
	var cars []*CarRecord
	for res.Next(c) {
		var cr CarRecord
		err = res.Decode(&cr)
		if err != nil {
			return nil, fmt.Errorf("cannot decode : %v,", err)
		}
		cars = append(cars, &cr)
	}
	return cars, nil
}

func (m *Mongo) UpdateCar(c context.Context, carID id.CarID, status carpb.CarStatus, car *CarUpdate) (*CarRecord, error) {
	objID, err := objid.FromID(carID)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}
	filter := bson.M{
		mgo.IDFieldName: objID,
	}
	if status != carpb.CarStatus_CS_NOT_SPECIFIED {
		filter[statusField] = status
	}
	u := bson.M{}
	if car.Status != carpb.CarStatus_CS_NOT_SPECIFIED {
		u[statusField] = car.Status
	}
	if car.Driver != nil {
		u[driverField] = car.Driver
	}
	if car.Position != nil {
		u[positionField] = car.Position
	}
	if car.UpdateTripID {
		u[tripIdField] = car.TripID.String()
	}

	res := m.col.FindOneAndUpdate(c, filter, mgo.Set(u), options.FindOneAndUpdate().SetReturnDocument(options.After))
	return convertSingleResult(res)
}

func convertSingleResult(res *mongo.SingleResult) (*CarRecord, error) {
	if err := res.Err(); err != nil {
		return nil, err
	}

	var cr CarRecord
	err := res.Decode(&cr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode: %v", err)
	}
	return &cr, nil
}
