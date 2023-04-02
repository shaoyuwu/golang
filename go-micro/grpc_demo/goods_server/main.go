package main

import (
	"context"
	"fmt"
	"goods_server/proto/goods"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Goods struct{}

func (t *Goods) AddGoods(c context.Context, req *goods.AddGoodsReq) (*goods.AddGoodsRes, error) {
	fmt.Println(req)
	return &goods.AddGoodsRes{
		Message: "增加数据成功",
		Success: true,
	}, nil
}

func (t *Goods) GetGoods(c context.Context, req *goods.GetGoodsReq) (*goods.GetGoodsRes, error) {

	var tempList []*goods.GoodsModel
	for i := 0; i < 10; i++ {
		tempList = append(tempList, &goods.GoodsModel{
			Title:   "golang入门到放弃" + strconv.Itoa(i),
			Price:   float64(i),
			Content: "golang基础" + strconv.Itoa(i),
		})
	}
	return &goods.GetGoodsRes{
		GoodsList: tempList,
	}, nil

}

func main() {
	//1.获取grpc对象
	grpcServer := grpc.NewServer()
	//2.注册服务
	goods.RegisterGoodsServer(grpcServer, new(Goods))

	//3.启动监听
	listenr, err := net.Listen("tcp", "127.0.0.1:8091")
	if err != nil {
		fmt.Println("启动监听报错:", err)
	}

	//4.关闭监听
	defer listenr.Close()

	//5.启动服务
	grpcServer.Serve(listenr)
}
