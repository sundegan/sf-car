package mgutil

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// Common field names.
const (
	IDFieldName        = "_id"
	UpdatedAtFieldName = "updatedat"
)

// IDField defines the object id field.
type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

// UpdatedAtField defines the UpdatedAt field.
type UpdatedAtField struct {
	UpdatedAt int64 `bson:"updatedat"`
}

// NewObjID generates a new object id.
var NewObjID = primitive.NewObjectID

// NewObjIDFormHex generates a object id from Hex string.
func NewObjIDFormHex(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		log.Fatalf("create ObjectID failed: %v", err)
	}
	return objID
}

// UpdatedAt returns a value suitable for UpdatedAt field.
var UpdatedAt = func() int64 {
	return time.Now().UnixNano()
}

// Set returns a $set update document.
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert returns a $setOnInsert update document.
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

// ZeroOrDoesNotExist generates a filter expression with
// field equal to zero or field does not exist.
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
