package main

import (
	"fmt"
	"gindemo/models"
	"gindemo/routers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func initMiddleware(c *gin.Context) {
	fmt.Println("中间件开始")
	//调用该请求的剩余处理程序
	c.Next()
	fmt.Println("中间件结束")

}
func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//自定义模板函数，注意要把这个函数放在加载函数模板之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	// r.GET("/test", initMiddleware, func(ctx *gin.Context) {
	// 	fmt.Println("这是一个首页")
	// 	ctx.String(200, "GIN首页")
	// })

	// //全局中间件
	// r.Use(initMiddleware) //可以配置多个，逗号间隔，使用后生效

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)
	r.Run(":8080")
}
