package trip

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	rentalpb "sfcar/rental/api/gen/v1"
)

// Service implements a trip service.
type Service struct {
	Logger *zap.Logger
	*rentalpb.UnimplementedTripServiceServer
}

// CreateTrip creates a trip.
func (s *Service) CreateTrip(ctx context.Context, request *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	s.Logger.Info("create trip", zap.String("account_id", request.AccountId))
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) mustEmbedUnimplementedTripServiceServer() {

}
