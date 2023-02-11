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
	*rentalpb.UnimplementedTripServiceServer
}

// ProfileManager defines the ACL for profile verification logic.
type ProfileManager interface {
	// Verify verifies that the account is eligible to create a trip.
	Verify(context.Context, id.AccountID) error
}

// CreateTrip creates a trip.
func (s *Service) CreateTrip(ctx context.Context, request *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	accountID, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// Verify driver's license identity.
	s.ProfileManager.Verify(ctx, accountID)
	s.Logger.Info("create trip", zap.String("account_id", accountID.String()))
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetTrip(ctx context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, nil
}

func (s *Service) GetTrips(ctx context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, nil
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
