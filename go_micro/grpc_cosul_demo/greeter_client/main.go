package main

import (
	"context"
	"fmt"
	"greeter_client/proto/greeter"
	"strconv"

	"github.com/hashicorp/consul/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//注册consul服务
	//1.初始化consul服务
	consulConfig := api.DefaultConfig()
	// consulConfig.Address = "127.0.0.1:8500"  consul默认地址

	//2.获取consul操作对象
	consulClient, _ := api.NewClient(consulConfig)

	//3.获取地址
	serviceEntry, _, _ := consulClient.Health().Service("HelloService", "test", false, nil)
	fmt.Println(serviceEntry[0].Service.Address)
	fmt.Println(serviceEntry[0].Service.Port)

	address := serviceEntry[0].Service.Address + ":" + strconv.Itoa(serviceEntry[0].Service.Port)
	//grpc客户端
	//连接服务器
	grpcClient, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务端失败:%s", err)
	}

	/*
		grpcClient, err := grpc.Dial(
			"consul://127.0.0.1:8500/HelloService",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
		if err != nil {
			fmt.Printf("grpc client连接报错:%v", err)
		}
	*/
	//2.注册客户端
	client := greeter.NewGreeterClient(grpcClient)

	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "golang",
	})
	fmt.Printf("%+v", res)
}
