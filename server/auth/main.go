package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	authpb "sfcar/auth/api/gen/v1"
	"sfcar/auth/auth"
	"sfcar/auth/dao"
	"sfcar/auth/wechat"
)

// 将auth服务注册到GRPC并启动auth GRPC服务
func main() {
	fmt.Println("start GRPC server...")
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed create logger: %v", err)
	}

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("failed listen at tcp:8081", zap.Error(err))
	}

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		logger.Fatal("connect to mondodb failed: %v", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wx2574ac10292f87b5",
			AppSecret: "176988fabd721c57829111c9d22a6199",
		},
		Mongo:  dao.NewMongo(mongoClient.Database("sfcar")),
		Logger: logger,
	})

	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("failed start server", zap.Error(err))
	}
}
