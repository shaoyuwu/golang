package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	dsn := "root:Wy!@tcp(172.23.102.55:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
		//SkipDefaultTransaction: true, //禁用事务
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}
}
