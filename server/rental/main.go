package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"sfcar/internal/auth_util"
	rentalpb "sfcar/rental/api/gen/v1"
	"sfcar/rental/trip"
)

// Register the auth service with GRPC and start the auth GRPC service.
func main() {
	fmt.Println("start Trip GRPC server...")
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed create logger: %v", err)
	}

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("failed listen at tcp:8082", zap.Error(err))
	}

	// Get the interceptor from auth_util package.
	in, err := auth_util.Interceptor("internal/auth_util/public.key")
	if err != nil {
		logger.Fatal("cannot create auth interceptor", zap.Error(err))
	}
	// Install the Token interceptor.
	s := grpc.NewServer(grpc.UnaryInterceptor(in))

	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})

	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("failed start Trip GRPC server", zap.Error(err))
	}
}
