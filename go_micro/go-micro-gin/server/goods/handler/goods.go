package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "goods/proto"
)

type Goods struct{}

func (e *Goods) AddGoods(ctx context.Context, req *pb.AddRequest, rsp *pb.AddResponse) error {
	log.Infof("Received Goods.Call request: %v", req)
	rsp.Message = "增加数据成功"
	rsp.Success = true
	return nil
}
