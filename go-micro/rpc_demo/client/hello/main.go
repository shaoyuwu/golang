package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1.用rpc.Dial和rpc微服务端建立连接
	conn, err := rpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//客户端退出时断开连接
	defer conn.Close()
	//2.调用远程函数

	var reply string
	err2 := conn.Call("hello.SayHello", " i am client", &reply)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	//4.获取微服务端返回的数据
	fmt.Printf("reply: %v\n", reply)
}
