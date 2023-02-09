package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	mongotest "sfcar/internal/mongo"
	"testing"
)

var mongodbURI string // mongodb connection URI

func TestMongo_ResolveAccountID(t *testing.T) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err != nil {
		t.Fatalf("connect to mondodb failed: %v", err)
	}
	client.Database("sfcar")
	m := NewMongo(client.Database("sfcar"))

	// Inserting initial data into mongodb.
	_, err = m.col.InsertMany(ctx, []interface{}{
		bson.M{
			"_id":     mongotest.NewObjID("63e26d0625d9b723e3f81901"),
			"open_id": "openid_1",
		},
		bson.M{
			"_id":     mongotest.NewObjID("63e26d0625d9b723e3f81902"),
			"open_id": "openid_2",
		},
	})
	if err != nil {
		log.Fatalf("cannot insert initial values: %v", err)
	}

	// This objID is used when new data are inserted.
	m.objID = mongotest.NewObjID("63e26d0625d9b723e3f81900")

	// Table-driven testing
	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user_1",
			openID: "openid_1",
			want:   "63e26d0625d9b723e3f81901",
		},
		{
			name:   "existing_user_2",
			openID: "openid_2",
			want:   "63e26d0625d9b723e3f81902",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			want:   "63e26d0625d9b723e3f81900",
		},
	}

	// Running the test case
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), c.openID)
			if err != nil {
				t.Errorf("failed resolve account id from %q: %v", c.openID, err)
			}
			if id != c.want {
				t.Errorf("resolve account id error: want: %q, got: %q", c.want, id)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotest.RunWithMongoInDocker(m, &mongodbURI))
}
