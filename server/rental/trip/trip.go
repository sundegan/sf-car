package trip

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sfcar/internal/auth_util"
	rentalpb "sfcar/rental/api/gen/v1"
)

// Service implements a trip service.
type Service struct {
	Logger *zap.Logger
	*rentalpb.UnimplementedTripServiceServer
}

// CreateTrip creates a trip.
func (s *Service) CreateTrip(ctx context.Context, request *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	accountID, err := auth_util.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// Get accountID from context.
	s.Logger.Info("create trip", zap.String("name:", request.Name), zap.String("account_id", accountID))
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) mustEmbedUnimplementedTripServiceServer() {

}
