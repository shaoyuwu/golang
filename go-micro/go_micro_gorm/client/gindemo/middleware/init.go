package middlerware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	//判断用户是否登录

	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)

	c.Set("username", "zhangsan")

	//定义一个goroutine统计日志，挡在中间件或handler中启动新的goroutine时，不能使用*gin.Context的上下文

	cCp := c.Copy()

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("done in path" + cCp.Request.URL.Path)
	}()
}
