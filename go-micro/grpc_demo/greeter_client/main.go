package main

import (
	"context"
	"fmt"
	"greeter_client/proto/greeter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//连接服务器

	grpcClient, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务端失败:%s", err)
	}
	//2.注册客户端
	client := greeter.NewGreeterClient(grpcClient)

	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "golang",
	})
	fmt.Printf("%+v", res)
}
