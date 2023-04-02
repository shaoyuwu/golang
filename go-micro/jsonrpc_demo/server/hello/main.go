package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 1.定义一个远程调用的方法
type Hello struct {
}

/*
1.方法只能有两个可序列化的参数，其中第二个是指针类型
req 表示获取用户传过来的数据
res 表示给客户端返回的数据
2.方法要返回一个error类型，同时必须是公开的方法

3.req和res的类型不能是:channel(类型)、func(函数)均不能进行序列化
*/
func (th Hello) SayHello(req string, res *string) error {
	fmt.Printf("req: %v\n", req)
	*res = "hello" + req
	return nil
}

func main() {
	//1.注册RPC服务,指定服务名称，不指定默认为结构体名称
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println("err:", err)
	}
	//2.监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println("err2:", err2)
	}

	//3.应用退出的时候关闭监听端口
	defer listener.Close()
	for {
		fmt.Println("开始建立连接")
		//4.建立连接、
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println("err3:", err3)
		}
		//5.绑定服务
		// rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		//也可使用这种方法启动
		//go jsonrpc.ServeConn(conn)  //这种写法会调用已经分装好的方法，及上面的写法
	}

}

//jsonrpc和默认rpc的区别：
/*
	之前rpc.ServeConn(conn)绑定服务
	jsonrpc中通过rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
*/
