package itying

import (
	"context"
	"gindemo/models"
	pb "gindemo/proto/goods"
	"strconv"

	"github.com/gin-gonic/gin"

	log "go-micro.dev/v4/logger"
	//注意
)

type GoodsController struct {
}

func (con GoodsController) Index(c *gin.Context) {
	/* 抽离
	//集成consul
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// Create service
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	srv.Init()
	*/

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := 5

	// Create client
	microClient := pb.NewGoodsService("goods", models.MicroClient)

	// Call service
	rsp, err := microClient.GetGoods(context.Background(), &pb.GetGoodsRequest{
		Page:     int64(page),
		PageSize: int64(pageSize),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Info(rsp)

	// time.Sleep(1 * time.Second)
	// c.String(200, "我是goods首页")
	c.JSON(200, gin.H{
		"result": rsp.GoodsList,
	})
}

/*
func (con GoodsController) Add(c *gin.Context) {
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
	microClient := pb.NewGoodsService("goods", models.MicroClient)

	// Call service
	rsp, err := microClient.AddGoods(context.Background(), &pb.AddGoodsRequest{
		Params: &pb.GoodsModel{
			Title: "增加商品", Price: 79.9, Content: "golang如梦",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Info(rsp)
	c.JSON(200, gin.H{
		"message": rsp.Message,
		"success": rsp.Success,
	})
}
*/
