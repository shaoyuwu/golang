package main

import (
	"greeter/handler"
	pb "greeter/proto"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "greeter"
	version = "latest"
)

func main() {
	//实例化consul
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Address("127.0.0.1:8080"), //指定微服务监听的地址
		micro.Name(service),
		micro.Version(version),
		//注册consul
		micro.Registry(consulReg),
	)
	srv.Init()

	// Register handler
	pb.RegisterGreeterHandler(srv.Server(), new(handler.Greeter))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
