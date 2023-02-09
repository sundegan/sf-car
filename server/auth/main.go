package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	authpb "sfcar/auth/api/gen/v1"
	"sfcar/auth/auth"
	"sfcar/auth/dao"
	"sfcar/auth/token"
	"sfcar/auth/wechat"
	"time"
)

// Register the auth service with GRPC and start the auth GRPC service.
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

	// Get the appsecret from the local file.
	appSecretFile, err := os.Open("auth/appsecret.txt")
	if err != nil {
		logger.Fatal("cannot open appsecret.txt file", zap.Error(err))
	}
	appSecretBytes, err := io.ReadAll(appSecretFile)
	if err != nil {
		logger.Fatal("cannot read appsecret.txt file", zap.Error(err))
	}
	appSecret := string(appSecretBytes)

	// Get the private key from the local file.
	keyFile, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("cannot open private key file", zap.Error(err))
	}
	keyBytes, err := io.ReadAll(keyFile)
	if err != nil {
		logger.Fatal("cannot read private key file", zap.Error(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		log.Fatalf("cannot parse private key: %v", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wx2574ac10292f87b5",
			AppSecret: appSecret,
		},
		Mongo:  dao.NewMongo(mongoClient.Database("sfcar")),
		Logger: logger,
		TokenGenerator: &token.JWTTokenGen{
			PrivateKey: privateKey,
			Issuer:     "sfcar/auth",
			IssuedAt:   time.Unix(1516239022, 0),
		},
		TokenExpire: 2 * time.Hour,
	})

	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("failed start server", zap.Error(err))
	}
}
