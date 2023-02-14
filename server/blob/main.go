package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	blobpb "sfcar/blob/api/gen/v1"
	"sfcar/blob/blob"
	"sfcar/blob/cos"
	"sfcar/blob/dao"
	"sfcar/internal/server"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}
	db := mongoClient.Database("sfcar")

	st, err := cos.NewService(
		"https://sfcar-1304689777.cos.ap-guangzhou.myqcloud.com",
		"AKIDoZG0c3MYeYFj2A97zhGfy98fAXqzveEc",
		"8B5oJar22pbGptlT0NaaVK2kKo9KLGz6",
	)
	if err != nil {
		logger.Fatal("cannot create cos service", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "Blob GRPC Server",
		Addr:   ":8083",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Storage: st,
				Mongo:   dao.NewMongo(db),
				Logger:  logger,
			})
		},
	}))
}
