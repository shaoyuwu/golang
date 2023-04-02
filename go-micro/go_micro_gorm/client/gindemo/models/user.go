package models

type User struct { //默认操作users表
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

//我们可以使用结构体中的自定义方法TableName改变结构体的默认表名称，如下：
//表示配置数据库表的名称
func (User) TableName() string {
	return "user"
}
