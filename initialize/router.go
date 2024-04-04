package initialize

import (
	"github.com/gin-gonic/gin"
	"gorm-example/global"
	"gorm-example/middleware"
	"gorm-example/router"
)

func Engin() *gin.Engine {

	gin.SetMode(global.Config.Runtime.Mode)

	engine := gin.New()

	engine.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		engine.Use(gin.Logger())
	}

	apiV1 := engine.Group(global.Config.System.PrefixUrl)

	apiV1.Use(middleware.TraceMiddleware())

	router.InitUserRouter(apiV1)

	return engine
}
