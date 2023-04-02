package main

import (
	"fmt"
	"proto_demo/proto/userService"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("hello golang")
	u := &userService.Userinfo{
		Username: "张三",
		Age:      20,
		Hobby:    []string{"吃饭", "睡觉"},
	}
	// fmt.Println(u)

	fmt.Println(u.GetType())
	fmt.Println(u.GetUsername())
	fmt.Println(u.GetHobby())
	//protobuf序列化
	data, _ := proto.Marshal(u)
	fmt.Println(data)

	//protobuf反序列化
	user := userService.Userinfo{}
	proto.Unmarshal(data, &user)
	fmt.Println(user)
	fmt.Println(user.GetHobby())
}
