package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo defines a mongo dao.
type Mongo struct {
	col *mongo.Collection
}

// NewMongo is used by external packages to initialize Mongo structs.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

// ResolveAccountID retrieves the AccountID from OpenID.
func (m *Mongo) ResolveAccountID(ctx context.Context, openID string) (string, error) {
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.M{
		"open_id": openID,
	}
	update := bson.M{
		"$setOnInsert": bson.M{
			"_id":     "63e26d0625d9b723e3f819ad",
			"open_id": openID,
		},
	}

	res := m.col.FindOneAndUpdate(
		ctx,
		filter,
		update,
		opts,
	)
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("failed to FindOneAndUpdate: %v", err)
	}

	var row struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("failed to Decode res: %v", err)
	}
	return row.ID.Hex(), nil
}