package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	authpb "sfcar/auth/api/gen/v1"
	rentalpb "sfcar/rental/api/gen/v1"
)

func main() {
	fmt.Println("start GRPC-Gateway server...")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// When GRPC-Gateway converts RPC to JSON,
	// the original field name is used and the
	// enumeration variable is represented by the enumeration value.
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:  true,
				UseEnumNumbers: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{},
		},
	))

	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register the authorization GRPC service to the GRPC-Gateway proxy server.
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:8081", // The address of the Auth GRPC service
		opts,             // Connection configuration
	)
	if err != nil {
		log.Fatalf("failed register Auth GPRC server to the GRPC-Gateway: %v", err)
	}

	// Register the Trip GRPC service to the GRPC-Gateway proxy server.
	err = rentalpb.RegisterTripServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:8082", // The address of the Trip GRPC service
		opts,             // Connection configuration
	)
	if err != nil {
		log.Fatalf("failed register Trip GPRC server to the GRPC-Gateway: %v", err)
	}

	// Start the GRPC-Gateway proxy server at port 8080.
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
