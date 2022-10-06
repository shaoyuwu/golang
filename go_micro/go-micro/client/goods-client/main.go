package main

import (
	"context"
	"time"

	pb "goods-client/proto"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry" //注意
)

var (
	service = "goods"
	version = "latest"
)

func main() {
	//集成consul
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// Create service
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	srv.Init()

	// Create client
	c := pb.NewGoodsService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.AddGoods(context.Background(), &pb.AddRequest{
			Title: "增加商品", Price: 79.9, Content: "golang如梦"})
		if err != nil {
			log.Fatal(err)
		}

		log.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}
