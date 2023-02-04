package main

import (
	"context"
	"log"

	service "helloworld/client/proto/gen/go"
	"google.golang.org/grpc"
)

func main() {
	// 连接GRPC服务
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 新建一个客户端
	client := service.NewGreeterClient(conn)
	req := &service.HelloRequest{
		Request: "world",
	}
	// 发送请求
	response, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("could not call SayHello: %v", err)
	}
	log.Printf("response: \n%s", response)
}