package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"sfcar/internal/server"
	rentalpb "sfcar/rental/api/gen/v1"
	"sfcar/rental/profile"
	prdao "sfcar/rental/profile/dao"
	"sfcar/rental/trip"
	trdao "sfcar/rental/trip/dao"
	"sfcar/rental/trip/impl/car"
	"sfcar/rental/trip/impl/poi"
	aclpr "sfcar/rental/trip/impl/profile"
)

// Register the auth service with GRPC and start the auth GRPC service.
func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed create zap logger: %v", err)
	}

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		logger.Fatal("connect to mondodb failed: %v", zap.Error(err))
	}
	db := mongoClient.Database("sfcar")

	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:          "Trip GRPC Server",
		Addr:          ":8082",
		PublicKeyFile: "internal/auth_util/public.key",
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger:         logger,
				CarManager:     &car.Manager{},
				ProfileManager: &aclpr.Manager{},
				POIManager:     &poi.Manager{},
				Mongo:          trdao.NewMongo(db),
			})
			rentalpb.RegisterProfileServiceServer(s, &profile.Service{
				Mongo:  prdao.NewMongo(db),
				Logger: logger,
			})
		},
		Logger: logger,
	})
	if err != nil {
		logger.Fatal("failed start Trip GRPC Server", zap.Error(err))
	}
}
