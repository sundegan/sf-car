package trip

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sfcar/internal/auth_util"
	"sfcar/internal/id"
	rentalpb "sfcar/rental/api/gen/v1"
	"sfcar/rental/trip/dao"
	"time"
)

// Service implements a trip service.
type Service struct {
	Logger         *zap.Logger
	Mongo          *dao.Mongo
	ProfileManager ProfileManager
	CarManager     CarManager
	POIManager     POIManager
	DistanceCalc   DistanceCalc
	*rentalpb.UnimplementedTripServiceServer
}

// ProfileManager defines the ACL for profile verification logic.
type ProfileManager interface {
	Verify(context.Context, id.AccountID) (id.IdentityID, error)
}

// CarManager defines the ACL for car management.
type CarManager interface {
	Verify(context.Context, id.CarID, *rentalpb.Location) error
	Unlock(context.Context, id.CarID) error
	Lock(ctx context.Context, carID id.CarID) error
}

// POIManager queries for the current landmark.
type POIManager interface {
	Resolve(ctx context.Context, location *rentalpb.Location) (string, error)
}

// DistanceCalc calculates distance between given points.
type DistanceCalc interface {
	DistanceKm(context.Context, *rentalpb.Location, *rentalpb.Location) (float64, error)
}

// CreateTrip creates a trip.
func (s *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aID, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// Check the passed arguments.
	if req.CarId == "" || req.Start == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}

	// Verify driver's license identity.
	iID, err := s.ProfileManager.Verify(ctx, aID)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	// Check the car status.
	carID := id.CarID(req.CarId)
	err = s.CarManager.Verify(ctx, carID, req.Start)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	// Create the trip record and write it to the database.
	poi, err := s.POIManager.Resolve(ctx, req.Start)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location", req.Start))
	}
	ls := &rentalpb.LocationStatus{
		Location:     req.Start,
		PoiName:      poi,
		TimestampSec: nowFunc(),
	}
	tr, err := s.Mongo.CreateTrip(ctx, &rentalpb.Trip{
		AccountId:  aID.String(),
		CarId:      carID.String(),
		IdentityId: iID.String(),
		Status:     rentalpb.TripStatus_IN_PROGRESS,
		Start:      ls,
		Current:    ls,
	})
	if err != nil {
		s.Logger.Warn("cannot create trip", zap.Error(err))
		return nil, status.Error(codes.AlreadyExists, "")
	}
	s.Logger.Info("create trip", zap.String("trip_id", tr.ID.String()))

	// Unlock the car in background.
	go func() {
		err = s.CarManager.Unlock(context.Background(), carID)
		if err != nil {
			s.Logger.Error("cannot unlock car", zap.Error(err))
		}
	}()

	return &rentalpb.TripEntity{
		Id:   tr.ID.Hex(),
		Trip: tr.Trip,
	}, nil
}

func (s *Service) GetTrip(ctx context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	aID, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tr, err := s.Mongo.GetTrip(ctx, id.TripID(req.Id), aID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return tr.Trip, nil
}

func (s *Service) GetTrips(ctx context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	aID, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	trips, err := s.Mongo.GetTrips(ctx, aID, req.Status)
	if err != nil {
		s.Logger.Error("cannot get trips", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	res := &rentalpb.GetTripsResponse{}
	for _, tr := range trips {
		res.Trips = append(res.Trips, &rentalpb.TripEntity{
			Id:   tr.ID.Hex(),
			Trip: tr.Trip,
		})
	}
	return res, nil
}

func (s *Service) UpdateTrip(ctx context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tid := id.TripID(req.Id)
	tr, err := s.Mongo.GetTrip(ctx, tid, aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}

	if tr.Trip.Status == rentalpb.TripStatus_FINISHED {
		return nil, status.Error(codes.FailedPrecondition, "cannot update a finished trip")
	}

	if tr.Trip.Current == nil {
		s.Logger.Error("trip without current set", zap.String("trip_id", tid.String()))
	}

	cur := tr.Trip.Current.Location
	if req.Current != nil {
		cur = req.Current
	}

	tr.Trip.Current = s.calcFeeCentAndKm(ctx, tr.Trip.Current, cur)

	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
		err := s.CarManager.Lock(ctx, id.CarID(tr.Trip.CarId))
		if err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "cannot lock car %v", err)
		}
	}
	err = s.Mongo.UpdateTrip(ctx, tid, aid, tr.UpdatedAt, tr.Trip)
	if err != nil {
		// try update again
		return nil, err
	}
	return tr.Trip, nil
}

// calcFeeCentAndKm Calculate the cost and distance traveled from the last location to the current location.
var nowFunc = func() int64 {
	return time.Now().Unix()
}

const (
	centsPerSec = 0.7
)

func (s *Service) calcFeeCentAndKm(ctx context.Context, last *rentalpb.LocationStatus, cur *rentalpb.Location) *rentalpb.LocationStatus {
	now := nowFunc()
	elapsedSec := float64(now - last.TimestampSec)

	dist, err := s.DistanceCalc.DistanceKm(ctx, last.Location, cur)
	if err != nil {
		s.Logger.Warn("cannot calcuate distance", zap.Error(err))
	}

	poi, err := s.POIManager.Resolve(ctx, cur)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location", cur))
	}

	return &rentalpb.LocationStatus{
		Location:     cur,
		FeeCent:      last.FeeCent + centsPerSec*elapsedSec,
		DrivenKm:     last.DrivenKm + dist,
		TimestampSec: now,
		PoiName:      poi,
	}
}

func (s *Service) mustEmbedUnimplementedTripServiceServer() {

}
