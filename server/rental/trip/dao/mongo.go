package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sfcar/internal/id"
	mgutil "sfcar/internal/mongo_util"
	rentalpb "sfcar/rental/api/gen/v1"
)

const (
	idField        = "_id"
	tripField      = "trip"
	accountIDField = tripField + ".accountid"
	statusField    = tripField + ".status"
)

// Mongo defines a mongo dao.
type Mongo struct {
	col *mongo.Collection
}

// NewMongo is used by external packages to initialize Mongo structs.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}

// TripRecord defines a trip record in mongo db.
type TripRecord struct {
	mgutil.IDField        `bson:"inline"`
	mgutil.UpdatedAtField `bson:"inline"`
	Trip                  *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(ctx context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	record := &TripRecord{
		Trip: trip,
	}
	record.ID = id.NewObjID()
	record.UpdatedAt = mgutil.UpdatedAt()

	_, err := m.col.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (m *Mongo) GetTrip(ctx context.Context, id id.TripID, accountID id.AccountID) (*TripRecord, error) {
	objID, err := primitive.ObjectIDFromHex(id.String())
	if err != nil {
		return nil, fmt.Errorf("invaild id: %v", err)
	}

	filter := bson.M{
		idField:        objID,
		accountIDField: accountID,
	}
	res := m.col.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		return nil, err
	}
	var trip TripRecord
	err = res.Decode(&trip)
	if err != nil {
		return nil, fmt.Errorf("cannot decode TripRecord: %v", err)
	}
	return &trip, nil
}

// GetTrips gets trips for the account by status.
// If status is not specified, gets all trips for the account.
func (m *Mongo) GetTrips(ctx context.Context, accountID id.AccountID, status rentalpb.TripStatus) ([]*TripRecord, error) {
	filter := bson.M{
		accountIDField: accountID.String(),
	}
	if status != rentalpb.TripStatus_TS_NOT_SPECIFIED {
		filter[statusField] = status
	}

	res, err := m.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var trips []*TripRecord
	for res.Next(ctx) {
		var trip TripRecord
		err := res.Decode(&trip)
		if err != nil {
			return nil, err
		}
		trips = append(trips, &trip)
	}
	return trips, nil
}

// UpdateTrip updates a trip.
func (m *Mongo) UpdateTrip(c context.Context, tid id.TripID, aid id.AccountID, updatedAt int64, trip *rentalpb.Trip) error {
	objID, err := id.ObjIDFromID(tid)
	if err != nil {
		return fmt.Errorf("invalid accountID: %v", err)
	}

	newUpdatedAt := mgutil.UpdatedAt()
	res, err := m.col.UpdateOne(c, bson.M{
		mgutil.IDFieldName:        objID,
		accountIDField:            aid.String(),
		mgutil.UpdatedAtFieldName: updatedAt,
	}, mgutil.Set(bson.M{
		tripField:                 trip,
		mgutil.UpdatedAtFieldName: newUpdatedAt,
	}))
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
