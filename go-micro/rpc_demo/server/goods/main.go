package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct{}

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

func (th Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	//1.执行增加
	fmt.Printf("req: %#v\n", req)
	//2.返回结果
	*res = AddGoodsRes{
		Success: true,
		Message: "增加数据成功",
	}
	return nil
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

func (th Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	//1.执行增加
	fmt.Printf("req: %#v\n", req)
	//2.返回结果
	*res = GetGoodsRes{
		Id:      12,
		Title:   "服务器获取的数据",
		Price:   24.5,
		Content: "我是服务器数据库获取的内容",
	}
	return nil
}

func main() {
	err := rpc.RegisterName("goods", new(Goods))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//2.监听端口
	listner, err2 := net.Listen("tcp", "127.0.0.1:8000")
	if err2 != nil {
		fmt.Println("err2:", err2)
	}
	//3.关闭端口监听
	defer listner.Close()

	for {
		//4.监听客户端连接
		fmt.Println("准备建立连接...")
		conn, err3 := listner.Accept()
		if err3 != nil {
			fmt.Printf("err3: %v\n", err3)
		}
		rpc.ServeConn(conn)
	}

}
