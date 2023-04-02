package routers

import (
	"gindemo/controllers/admin"
	middlerware "gindemo/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//配置中间件
	adminRouters := r.Group("/admin", middlerware.InitMiddleware)
	{
		adminRouters.GET("/", admin.IndexController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
		adminRouters.GET("/article/delete", admin.ArticleController{}.Delete)
	}
}
