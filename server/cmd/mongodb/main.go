package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	// connect mongodb database and creat a mongodb client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		log.Fatalf("failed to connect mongodb: %v", err)
	}
	col := client.Database("sfcar").Collection("account")
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	insertRows(ctx, col)
	findRows(ctx, col)
}

func insertRows(ctx context.Context, col *mongo.Collection) {
	res, err := col.InsertMany(ctx, []interface{}{
		bson.M{
			"open_id": "123",
		},
		bson.M{
			"open_id": "456",
		},
	})
	if err != nil {
		log.Fatalf("insert rows failed: %v", err)
	}
	fmt.Printf("%+v\n", res)
}

func findRows(ctx context.Context, col *mongo.Collection) {
	res := col.FindOne(ctx, bson.M{
		"open_id": "123",
	})
	fmt.Printf("%+v\n", res)
	var row struct {
		ID     primitive.ObjectID `bson:"_id"`
		OpenID string             `bson:"open_id"`
	}
	err := res.Decode(&row)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", row)
}
