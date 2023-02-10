package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"sfcar/internal/server"
	rentalpb "sfcar/rental/api/gen/v1"
	"sfcar/rental/trip"
)

// Register the auth service with GRPC and start the auth GRPC service.
func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed create zap logger: %v", err)
	}

	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:          "Trip GRPC Server",
		Addr:          ":8082",
		PublicKeyFile: "internal/auth_util/public.key",
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger: logger,
			})
		},
		Logger: logger,
	})
	if err != nil {
		logger.Fatal("failed start Trip GRPC Server", zap.Error(err))
	}
}
