package handler

import (
	"context"

	"goods/models"
	pb "goods/proto/goods"
)

type Goods struct{}

func (e *Goods) GetGoods(ctx context.Context, req *pb.GetGoodsRequest, rsp *pb.GetGoodsResponse) error {
	//获取page和pagesize
	page := req.Page
	pageSize := req.PageSize
	//查询数据库
	goodsList := []models.Goods{}
	models.DB.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&goodsList)
	var tempList []*pb.GoodsModel
	for _, v := range goodsList {
		tempList = append(tempList, &pb.GoodsModel{
			Title:    v.Title,
			SubTitle: v.SubTitle,
			GoodsSn:  v.GoodsSn,
			Price:    v.Price,
			Sort:     int64(v.Sort),
			Status:   int32(v.Status),
			AddTime:  int64(v.AddTime),
		})
	}

	rsp.GoodsList = tempList
	return nil
}
