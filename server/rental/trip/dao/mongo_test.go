package dao

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
	"os"
	"sfcar/internal/id"
	mgutil "sfcar/internal/mongo_util"
	rentalpb "sfcar/rental/api/gen/v1"
	"testing"
)

func TestMongo_CreateTrip(t *testing.T) {
	ctx := context.Background()
	client, err := mgutil.NewClient(ctx)
	if err != nil {
		t.Fatalf("connect to mondodb failed: %v", err)
	}
	db := client.Database("sfcar")
	err = mgutil.SetupIndexes(ctx, db)
	if err != nil {
		t.Fatalf("cannot setup index: %v", err)
	}
	m := NewMongo(db)

	cases := []struct {
		name       string
		tripID     string
		accountID  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finished",
			tripID:     "5f8132eb00714bf62948905c",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "another_finished",
			tripID:     "5f8132eb00714bf62948905d",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "in_progress",
			tripID:     "5f8132eb00714bf62948905e",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			name:       "another_in_progress",
			tripID:     "5f8132eb00714bf62948905f",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			name:       "in_progress_by_another_account",
			tripID:     "5f8132eb00714bf629489060",
			accountID:  "account2",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	for _, cc := range cases {
		id.NewObjID = func() primitive.ObjectID {
			return id.NewObjIDFormHex(cc.tripID)
		}
		tr, err := m.CreateTrip(ctx, &rentalpb.Trip{
			AccountId: cc.accountID,
			Status:    cc.tripStatus,
		})
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: error expected; got none", cc.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: error creating trip: %v", cc.name, err)
			continue
		}
		if tr.ID.Hex() != cc.tripID {
			t.Errorf("%s: incorrect trip id; want: %q; got: %q",
				cc.name, cc.tripID, tr.ID.Hex())
		}
	}
}

func TestMongo_GetTrip(t *testing.T) {
	ctx := context.Background()
	client, err := mgutil.NewClient(ctx)
	if err != nil {
		t.Fatalf("connect to mondodb failed: %v", err)
	}
	m := NewMongo(client.Database("sfcar"))
	id.NewObjID = primitive.NewObjectID
	trip, err := m.CreateTrip(ctx, &rentalpb.Trip{
		AccountId: "account1",
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			PoiName: "start",
		},
		Current: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  31,
				Longitude: 120,
			},
			FeeCent:  10,
			DrivenKm: 10,
			PoiName:  "current",
		},
		End: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  32,
				Longitude: 120,
			},
			FeeCent:  20,
			DrivenKm: 20,
			PoiName:  "end",
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Errorf("cannot create trip: %v", err)
	}

	got, err := m.GetTrip(ctx, id.ToTripID(trip.ID), "account1")
	if err != nil {
		t.Errorf("cannot get trip: %v", err)
	}

	if diff := cmp.Diff(trip, got, protocmp.Transform()); diff != "" {
		t.Errorf("result differs; -want +got: %s", diff)
	}
}

func TestMongo_GetTrips(t *testing.T) {
	rows := []struct {
		id        string
		accountID string
		status    rentalpb.TripStatus
	}{
		{
			id:        "5f8132eb10714bf629489051",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "5f8132eb10714bf629489052",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "5f8132eb10714bf629489053",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "5f8132eb10714bf629489054",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			id:        "5f8132eb10714bf629489055",
			accountID: "account_id_for_get_trips_1",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	c := context.Background()
	client, err := mgutil.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(client.Database("sfcar"))

	for _, r := range rows {
		id.NewObjID = func() primitive.ObjectID {
			return id.NewObjIDFormHex(r.id)
		}
		_, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: r.accountID,
			Status:    r.status,
		})
		if err != nil {
			t.Fatalf("cannot create rows: %v", err)
		}
	}

	cases := []struct {
		name       string
		accountID  string
		status     rentalpb.TripStatus
		wantCount  int
		wantOnlyID string
	}{
		{
			name:      "get_all",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCount: 4,
		},
		{
			name:       "get_in_progress",
			accountID:  "account_id_for_get_trips",
			status:     rentalpb.TripStatus_IN_PROGRESS,
			wantCount:  1,
			wantOnlyID: "5f8132eb10714bf629489054",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			res, err := m.GetTrips(context.Background(),
				id.AccountID(cc.accountID),
				cc.status)
			if err != nil {
				t.Errorf("cannot get trips: %v", err)
			}

			if cc.wantCount != len(res) {
				t.Errorf("incorrect result count; want: %d, got: %d",
					cc.wantCount, len(res))
			}

			if cc.wantOnlyID != "" && len(res) > 0 {
				if cc.wantOnlyID != res[0].ID.Hex() {
					t.Errorf("only_id incorrect; want: %q, got %q",
						cc.wantOnlyID, res[0].ID.Hex())
				}
			}
		})
	}
}

func TestUpdateTrip(t *testing.T) {
	c := context.Background()
	client, err := mgutil.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(client.Database("sfcar"))
	tid := id.TripID("5f8132eb12714bf629489054")
	aid := id.AccountID("account_for_update")

	id.NewObjID = func() primitive.ObjectID {
		return id.NewObjIDFormHex(tid.String())
	}
	var now int64 = 10000
	mgutil.UpdatedAt = func() int64 {
		return now
	}

	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi",
		},
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}
	if tr.UpdatedAt != 10000 {
		t.Fatalf("wrong updatedat; want: 10000, got: %d", tr.UpdatedAt)
	}

	update := &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi_updated",
		},
	}
	cases := []struct {
		name          string
		now           int64
		withUpdatedAt int64
		wantErr       bool
	}{
		{
			name:          "normal_update",
			now:           20000,
			withUpdatedAt: 10000,
		},
		{
			name:          "update_with_stale_timestamp",
			now:           30000,
			withUpdatedAt: 10000,
			wantErr:       true,
		},
		{
			name:          "update_with_refetch",
			now:           40000,
			withUpdatedAt: 20000,
		},
	}

	for _, cc := range cases {
		now = cc.now
		err := m.UpdateTrip(c, tid, aid, cc.withUpdatedAt, update)
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want error; got none", cc.name)
			} else {
				continue
			}
		} else {
			if err != nil {
				t.Errorf("%s: cannot update: %v", cc.name, err)
			}
		}
		updatedTrip, err := m.GetTrip(c, tid, aid)
		if err != nil {
			t.Errorf("%s: cannot get trip after udpate: %v", cc.name, err)
		}
		if cc.now != updatedTrip.UpdatedAt {
			t.Errorf("%s: incorrect updatedat: want %d, got %d",
				cc.name, cc.now, updatedTrip.UpdatedAt)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(mgutil.RunWithMongoInDocker(m))
}
