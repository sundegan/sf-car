package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	authpb "sfcar/auth/api/gen/v1"
	"sfcar/auth/auth"
	"sfcar/auth/dao"
	"sfcar/auth/token"
	"sfcar/auth/wechat"
	"sfcar/internal/server"
	"time"
)

// Register the auth service with GRPC and start the auth GRPC service.
func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed create zap logger: %v", err)
	}

	// Create a MongoDB client.
	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		logger.Fatal("connect to mondodb failed: %v", zap.Error(err))
	}

	// Get the AppS ecret from the local file.
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

	err = server.RunGRPCServer(&server.GRPCConfig{
		Name: "Auth GRPC Server",
		Addr: ":8081",
		RegisterFunc: func(s *grpc.Server) {
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
					IssuedAt:   time.Now(),
				},
				TokenExpire: 2 * time.Hour,
			})
		},
		Logger: logger,
	})
	if err != nil {
		logger.Fatal("failed start Auth GRPC Server", zap.Error(err))
	}
}
