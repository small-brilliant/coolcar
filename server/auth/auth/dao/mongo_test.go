package dao

import (
	"context"
	mongotesting "coolcar/shared/mongo/testing"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string

func TestResolverAccountID(t *testing.T) {
	c := context.Background()
	s := mongoURI
	mc, err := mongo.Connect(c, options.Client().ApplyURI(s))
	if err != nil {
		t.Fatalf("cannot connection db: %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	m.col.InsertMany(c, []interface{}{
		bson.M{
			"_id":     mustObjID("621ecd8d42ffa4ec7cf1a22c"),
			"open_id": "open_1",
		},
		bson.M{
			"_id":     mustObjID("621ecd8d42ffa4ec7cf1a23c"),
			"open_id": "open_2",
		},
		bson.M{
			"_id":     mustObjID("621ecd8d42ffa4ec7cf1a24c"),
			"open_id": "open_3",
		},
	})
	// m.newObjID = func() primitive.ObjectID {
	// 	return mustObjID("621ecd8d42ffa4ec7cf1a26c")
	// }
	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "open_1",
			want:   "621ecd8d42ffa4ec7cf1a22c",
		},
		{
			name:   "another_existing_user",
			openID: "open_2",
			want:   "621ecd8d42ffa4ec7cf1a23c",
		},
		{
			name:   "new_user",
			openID: "open_3",
			want:   "621ecd8d42ffa4ec7cf1a24c",
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolverAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Errorf("faild resolved account id %s: %v", cc.openID, err)
			}
			if id != cc.want {
				t.Errorf("resolve account id : want: %q, got: %q", cc.want, id)
			}
		})
	}
	id, err := m.ResolverAccountID(c, "123")
	if err != nil {
		t.Errorf("faild resolved account id 123: %v", err)
	} else {
		want := "621ecd8d42ffa4ec7cf1a22c"
		if id != want {
			t.Errorf("resolve account id : want: %q, got: %q", want, id)
		}
	}

}

func mustObjID(hex string) primitive.ObjectID {
	oi, _ := primitive.ObjectIDFromHex(hex)
	return oi
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
