package admin

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (th BaseController) success(c *gin.Context) {
	c.String(200, "成功")
}

func (th BaseController) error(c *gin.Context) {
	c.String(200, "失败")
}
