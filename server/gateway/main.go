package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	authpb "sfcar/auth/api/gen/v1"
	rentalpb "sfcar/rental/api/gen/v1"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed create zap logger: %v", err)
	}

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
	grpcServers := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "Auth GRPC Server",
			addr:         "localhost:8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "Trip GRPC Server",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}

	for _, s := range grpcServers {
		err := s.registerFunc(
			ctx,
			mux,
			s.addr,
			opts,
		)
		if err != nil {
			logger.Sugar().Fatalf("cannot register %s to the GRPC-Gateway: %v", s.name, err)
		} else {
			logger.Sugar().Infof("Register the %s to the GRPC-gateway service", s.name)
		}
	}

	// Start the GRPC-Gateway proxy server at port 8080.
	addr := ":8080"
	logger.Sugar().Infof("GRPC-Gateway Server started at %s", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Sugar().Fatal(err)
	}
}
