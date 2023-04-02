package main

import (
	"context"
	"fmt"
	"greeter_server/proto/greeter"
	"net"

	"google.golang.org/grpc"
)

//1、定义远程调用的结构体和方法 我们这个结构体需要实现GreeterServer的接口

type Hello struct{}

func (th *Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{
		Message: "你好" + req.Name,
	}, nil
}

func main() {
	//注册grpc服务
	//1.初始化一个grpc对象
	grpcServer := grpc.NewServer()
	//2.注册服务
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	//3.设置监听，指定IP PORT
	listenr, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	//4.关闭服务
	defer listenr.Close()
	//5.启动服务
	grpcServer.Serve(listenr)
}
