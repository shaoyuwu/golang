package routers

import (
	"gindemo/controllers/itying"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)

		defaultRouters.GET("/news", itying.DefaultController{}.News)
		defaultRouters.GET("/goods", itying.GoodsController{}.Index)
		// defaultRouters.GET("/goods/add", itying.GoodsController{}.Add)
	}
}
