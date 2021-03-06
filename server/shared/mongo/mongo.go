package mgo

import (
	"coolcar/shared/mongo/objid"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var NewObjID = primitive.NewObjectID

func NewObjIDWithValue(id fmt.Stringer) {
	NewObjID = func() primitive.ObjectID {
		return objid.MustFromID(id)
	}
}

var UpdateAt = func() int64 {
	return time.Now().UnixNano()
}

const (
	IDFieldName        = "_id"
	UpdatedAtFieldName = "updatedat"
)

func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

type UpdatedField struct {
	UpdatedAt int64 `bson:"updatedat"`
}
type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

// 更新的条件
func ZeroOrDoesNotExist(field string, zero interface{}) bson.M {
	return bson.M{
		"$or": []bson.M{
			{
				field: zero,
			},
			{
				field: bson.M{
					"$exists": false,
				},
			},
		},
	}
}
