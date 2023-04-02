package handler

import (
	"context"
	"strconv"

	log "go-micro.dev/v4/logger"

	pb "goods/proto"
)

type Goods struct{}

func (e *Goods) AddGoods(ctx context.Context, req *pb.AddGoodsRequest, rsp *pb.AddGoodsResponse) error {
	log.Infof("Received Goods.Call request: %v", req)
	rsp.Message = "增加数据成功"
	rsp.Success = true
	return nil
}

func (e *Goods) GetGoods(ctx context.Context, req *pb.GetGoodsRequest, rsp *pb.GetGoodsResponse) error {

	var tempList []*pb.GoodsModel
	for i := 1; i < 10; i++ {
		tempList = append(tempList, &pb.GoodsModel{
			Title:   "这是第" + strconv.Itoa(i) + "条数据",
			Price:   79.9,
			Content: "这是第" + strconv.Itoa(i) + "条内容",
		})
	}
	rsp.GoodsList = tempList
	return nil
}
