package main

import (
	"context"
	"fmt"
	"goods_client/proto/goods"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//连接服务器

	grpcClient, err := grpc.Dial("127.0.0.1:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务端失败:%s", err)
	}
	//2.注册客户端
	client := goods.NewGoodsClient(grpcClient)

	//3.增加商品
	res1, err := client.AddGoods(context.Background(), &goods.AddGoodsReq{
		Goods: &goods.GoodsModel{
			Title:   "grpc框架",
			Price:   20,
			Content: "gprc need protobuf",
		},
	})
	fmt.Println(res1.Message)
	fmt.Println(res1.Success)

	//获取商品详情

	res2, _ := client.GetGoods(context.Background(), &goods.GetGoodsReq{})
	fmt.Printf("%v", res2)
	fmt.Printf("%v", res2.GoodsList)

}
