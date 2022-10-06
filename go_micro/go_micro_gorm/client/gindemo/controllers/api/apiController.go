package api

import "github.com/gin-gonic/gin"

type ApiController struct {
}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "我是api接口")
}

func (con ApiController) List(c *gin.Context) {
	c.String(200, "我是api/userlist接口")
}

func (con ApiController) Plist(c *gin.Context) {
	c.String(200, "我是api/plist接口")
}

func (con ApiController) Cart(c *gin.Context) {
	c.String(200, "我是api/cart接口")
}
