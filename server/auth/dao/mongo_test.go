package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	mongotesting "sfcar/internal/mongo"
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
	id, err := m.ResolveAccountID(ctx, "123")
	if err != nil {
		t.Fatalf("failed resolve account id from openid: %v", err)
	} else {
		want := "63e26d0625d9b723e3f819ad"
		if id != want {
			t.Errorf("resolve account id error: want: %q, got: %q", want, id)
		}
	}
}

func TestMain(m *testing.M) {

	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongodbURI))
}
