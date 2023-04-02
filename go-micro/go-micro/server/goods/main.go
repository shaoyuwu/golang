package main

import (
	"goods/handler"
	pb "goods/proto"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "goods"
	version = "latest"
)

func main() {
	//注册consul
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)
	srv.Init()

	// Register handler
	pb.RegisterGoodsHandler(srv.Server(), new(handler.Goods))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
