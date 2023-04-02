package main

import (
	"context"
	"time"

	pb "greeter-client/proto"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "greeter" //这个名称需要和微服务中的一样
	version = "latest"
)

func main() {
	//集成consul
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	srv.Init()

	// Create client
	c := pb.NewGreeterService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.Call(context.Background(), &pb.CallRequest{Name: "John"})
		if err != nil {
			log.Fatal(err)
		}

		log.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}
