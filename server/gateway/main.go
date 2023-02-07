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
)

func main() {
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
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:8081", // GRPC服务的地址
		opts,             // 连接配置，只能是切片类型
	)
	if err != nil {
		log.Fatalf("failed register auth GPRC server to the GRPC-Gateway: %v", err)
	}

	// 启动GRPC-Gateway代理服务器,地址为8080
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
