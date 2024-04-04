package router

import (
	"github.com/gin-gonic/gin"
	"gorm-example/controller"
)

func InitUserRouter(r *gin.RouterGroup) {

	user := r.Group("/user")

	user.GET("/join", controller.JoinFunc)
	user.GET("/preload", controller.PreloadFunc)
	user.GET("/preloads", controller.PreloadsFunc)
}
