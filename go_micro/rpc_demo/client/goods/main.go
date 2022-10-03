package main

import (
	"fmt"
	"net/rpc"
)

//AddGoods参数对应的结构体
type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsRes struct {
	Success bool
	Message string
}

//GetGoods参数对应的结构体
type GetGoodsReq struct {
	Id int
}

type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func main() {
	//1.用rpc.Dial和rpc微服务端建立连接
	conn, err := rpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//客户端退出时断开连接
	defer conn.Close()

	//2.调用远程函数

	var reply AddGoodsRes
	err2 := conn.Call("goods.AddGoods", AddGoodsReq{
		Id:      10,
		Title:   "商品标题",
		Price:   23.5,
		Content: "商品详情",
	}, &reply)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	//4.获取微服务端返回的数据
	fmt.Printf("reply: %v\n", reply)

	// 5.调用远程GetGoods函数
	var goodsData GetGoodsRes
	err3 := conn.Call("goods.GetGoods", GetGoodsReq{
		Id: 12,
	}, &goodsData)
	if err3 != nil {
		fmt.Println("err3:", err3)
	}
	fmt.Printf("%#v", goodsData)

}
