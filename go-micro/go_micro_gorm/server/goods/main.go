package main

import (
	"goods/handler"
	pb "goods/proto/goods"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "goods"
	version = "latest"
)

func main() {
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
