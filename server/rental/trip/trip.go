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
)

// Service implements a trip service.
type Service struct {
	Logger         *zap.Logger
	Mongo          *dao.Mongo
	ProfileManager ProfileManager
	CarManager     CarManager
	POIManager     POIManager
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
}

// POIManager queries for the current landmark.
type POIManager interface {
	Resolve(ctx context.Context, location *rentalpb.Location) (string, error)
}

// CreateTrip creates a trip.
func (s *Service) CreateTrip(ctx context.Context, request *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	accountID, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// Verify driver's license identity.
	identityID, err := s.ProfileManager.Verify(ctx, accountID)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	// Check the car status.
	carID := id.CarID(request.CarId)
	err = s.CarManager.Verify(ctx, carID, request.Start)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	// Create the trip record and write it to the database.
	poi, err := s.POIManager.Resolve(ctx, request.Start)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location", request.Start))
	}
	ls := &rentalpb.LocationStatus{
		Location: request.Start,
		PoiName:  poi,
	}
	tr, err := s.Mongo.CreateTrip(ctx, &rentalpb.Trip{
		AccountId:  accountID.String(),
		CarId:      carID.String(),
		IdentityId: identityID.String(),
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
	accountID, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tr, err := s.Mongo.GetTrip(ctx, id.TripID(req.Id), accountID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return tr.Trip, nil
}

func (s *Service) GetTrips(ctx context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	accountID, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	trips, err := s.Mongo.GetTrips(ctx, accountID, req.Status)
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
		return nil, status.Errorf(codes.Unauthenticated, "")
	}
	trip, err := s.Mongo.GetTrip(ctx, id.TripID(req.Id), aid)
	if req.Current != nil {
		trip.Trip.Current = s.calcFeeCentAndKm(trip.Trip, req.Current)
	}
	if req.EndTrip {
		trip.Trip.End = trip.Trip.Current
		trip.Trip.Status = rentalpb.TripStatus_FINISHED
	}
	return nil, nil
}

func (s *Service) calcFeeCentAndKm(trip *rentalpb.Trip, cur *rentalpb.Location) *rentalpb.LocationStatus {
	// Calculate the cost and distance
	return nil
}

func (s *Service) mustEmbedUnimplementedTripServiceServer() {

}
