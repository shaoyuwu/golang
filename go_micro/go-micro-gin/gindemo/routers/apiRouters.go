package routers

import (
	"gindemo/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userlist", api.ApiController{}.List)
		apiRouters.GET("/plist", api.ApiController{}.Plist)
		apiRouters.GET("/cart", api.ApiController{}.Cart)
	}
}
