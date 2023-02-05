package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	service "helloworld/proto/gen/go"
)

// server is used to implement the Greeter service.
type server struct {
	service.UnimplementedGreeterServer
}

// SayHello implements the SayHello method of the Greeter service.
func (s *server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	response := &service.HelloResponse{
		Response: "hello, " + req.Request,
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
	service.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGrpcGateway() {
	fmt.Println("start GRPC-Gateway server...")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// 注意：runtime是gprc-gateway组件下的,不是Go源码中的runtime
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 将具体的GPRC服务注册到GRPC-Gateway代理服务器
	err := service.RegisterGreeterHandlerFromEndpoint(
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

// grpc server start...
// start GRPC-Gateway server...
