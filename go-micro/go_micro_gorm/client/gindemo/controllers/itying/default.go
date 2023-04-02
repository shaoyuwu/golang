package itying

import (
	"fmt"
	"gindemo/models"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(c *gin.Context) {
	fmt.Println(models.UnixToTime(1629788564))
	// c.String(200, "首页")
	c.HTML(200, "default/index.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788418,
	})
}

func (con DefaultController) News(c *gin.Context) {
	c.String(200, "新闻")
}
