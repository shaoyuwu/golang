package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func (con IndexController) Index(c *gin.Context) {

	username, _ := c.Get("username")
	fmt.Printf("%+v\n", username)
	// con.success(c)
	s, ok := username.(string)
	if ok {
		c.String(200, "我是admin接口首页"+s)
	} else {
		c.String(200, "我是admin接口首页")
	}

}
