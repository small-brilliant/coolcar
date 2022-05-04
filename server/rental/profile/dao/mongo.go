package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	accountIdField      = "accoundid"
	profileField        = "profile"
	identityStatusField = profileField + ".identitystatus"
	photoBlobIDField    = "photoblobid"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("profile"),
	}
}

type ProfileRecord struct {
	AccountID   string            `bson:"accoundid"`
	Profile     *rentalpb.Profile `bson:"profile"`
	PhotoBlobID string            `bson:"photoblobid"`
}

func (m *Mongo) GetProfile(c context.Context, aid id.AccountID) (*ProfileRecord, error) {
	res := m.col.FindOne(c, byAccountID(aid))
	if err := res.Err(); err != nil {
		return nil, err
	}
	var pr ProfileRecord
	err := res.Decode(&pr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode profile record:%v", err)
	}
	return &pr, nil
}

func (m *Mongo) UpdateProfilePhoto(c context.Context, aid id.AccountID, bid id.BlobID) error {
	_, err := m.col.UpdateOne(c, bson.M{
		accountIdField: aid.String(),
	}, mgo.Set(bson.M{
		photoBlobIDField: bid.String(),
	}), options.Update().SetUpsert(true))
	return err
}

// 在UpdateProfilePhoto的时候就已经插入了account了

func (m *Mongo) UpdateProfile(c context.Context, aid id.AccountID, prevStatus rentalpb.IdentityStatus, p *rentalpb.Profile) error {
	filter := bson.M{
		identityStatusField: prevStatus,
	}
	if prevStatus == rentalpb.IdentityStatus_UNSUBMITTED {
		filter = mgo.ZeroOrDoesNotExist(identityStatusField, prevStatus)
	}
	filter[accountIdField] = aid.String()
	_, err := m.col.UpdateOne(c, filter, mgo.Set(bson.M{
		accountIdField: aid.String(),
		profileField:   p,
	}), options.Update().SetUpsert(true))
	return err
}
func byAccountID(aid id.AccountID) bson.M {
	return bson.M{
		accountIdField: aid.String(),
	}
}
