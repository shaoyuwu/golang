package admin

import "github.com/gin-gonic/gin"

type ArticleController struct {
	BaseController
}

func (con ArticleController) Index(c *gin.Context) {
	// c.String(200, "admin/article接口==")
	con.success(c)
}

func (con ArticleController) Edit(c *gin.Context) {
	c.String(200, "admin/article/edit接口==")
}

func (con ArticleController) Delete(c *gin.Context) {
	c.String(200, "admin/article/delete接口")
}
