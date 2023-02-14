package id

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// AccountID defines account id object.
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

// TripID defines account id object.
type TripID string

func (t TripID) String() string {
	return string(t)
}

// IdentityID defines identity id object.
type IdentityID string

func (i IdentityID) String() string {
	return string(i)
}

// CarID defines car id object.
type CarID string

func (c CarID) String() string {
	return string(c)
}

// BlobID defines blob id object.
type BlobID string

func (b BlobID) String() string {
	return string(b)
}

// ObjIDFromID converts an id to mongodb object id.
func ObjIDFromID(id fmt.Stringer) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

// ToAccountID converts a mongodb object id to account id.
func ToAccountID(objID primitive.ObjectID) AccountID {
	return AccountID(objID.Hex())
}

// ToTripID converts a mongodb object id to trip id.
func ToTripID(objID primitive.ObjectID) TripID {
	return TripID(objID.Hex())
}

// NewObjID generates a new object id.
// By default, it generates id based on the current time.
// If you need to, you can change the way of the id generation.
var NewObjID = primitive.NewObjectID

// NewObjIDFormHex generates an object id from Hex string.
func NewObjIDFormHex(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		log.Fatalf("create ObjectID failed: %v", err)
	}
	return objID
}

// ChangeToFromHex change the way of objectID generation to NewObjIDFormHex().
func ChangeToFromHex(hex string) {
	NewObjID = func() primitive.ObjectID {
		return NewObjIDFormHex(hex)
	}
}
