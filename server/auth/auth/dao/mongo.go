package dao

import (
	"context"
	mgo "coolcar/shared/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	col    *mongo.Collection
	NewObj func() primitive.ObjectID
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:    db.Collection("account"),
		NewObj: primitive.NewObjectID,
	}
}

func (m *Mongo) ResolverAccountID(c context.Context, openID string) (string, error) {

	insertedID := m.NewObj()
	res := m.col.FindOneAndUpdate(c, bson.M{
		"open_id": openID,
	}, mgo.SetOnInsert(bson.M{
		"_id":     insertedID,
		"open_id": openID,
	}), options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))

	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %v", err)
	}
	var row mgo.ObjID

	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %v", err)
	}
	return row.ID.Hex(), nil
}
