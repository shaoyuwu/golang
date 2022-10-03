package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//1.用net.Dial和jsonrpc微服务端建立连接
	// conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//客户端退出时断开连接
	defer conn.Close()

	//建立基于json编解码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	//2.调用远程函数

	var reply string
	err2 := client.Call("hello.SayHello", "我是客户端", &reply)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	//4.获取微服务端返回的数据
	fmt.Printf("reply: %v\n", reply)
}

/*
把默认的rpc改为jsonrpc
1、rpc.Dial需要调换成net.Dial
2、增加建立基于json编解码的rpc服务，client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
3、conn.Call 需要改为client.Call
*/
