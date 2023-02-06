package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	trippb "sfcar/proto/gen/go"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server struct {
	trippb.UnimplementedTripServiceServer
}

func (s *server) GetTrip(ctx context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	trip := trippb.Trip{
		Start:       "abc",
		End:         "def",
		DurationSec: 3600,
		FeeCent:     1000,
		StartPos: &trippb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		EndPos: &trippb.Location{
			Latitude:  35,
			Longitude: 120,
		},
		Status: trippb.TripStatus_IN_PROGRESS,
	}

	response := &trippb.GetTripResponse{
		Id:   req.TripId,
		Trip: &trip,
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go startGrpcGateway()

	fmt.Println("grpc server start...")
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGrpcGateway() {
	fmt.Println("start GRPC-Gateway server...")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	// GRPC-Gateway将RPC转为JSON时，使用原始字段名、枚举变量使用枚举值表示
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
	// 将具体的GPRC服务注册到GRPC-Gateway代理服务器
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:8081", // GRPC服务的地址
		opts,             // 连接配置，只能是切片
	)
	if err != nil {
		fmt.Println(err)
	}

	// 启动GRPC-Gateway代理服务器,地址为8080
	http.ListenAndServe(":8080", mux)
}
