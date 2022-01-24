package initialize

import (
	"github.com/gin-gonic/gin"
	"personFrame/middleware"
	"personFrame/pkg/common"
	"personFrame/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	common.GlobalLog.Info("use middleware logger")
	v1Router := router.RouterGroupApp.RouterGroupV1

	//注册公共Api不鉴权
	PublicGroup := Router.RouterGroup.Group("/api")
	{
		v1Router.InitBaseRouter(PublicGroup)
	}

	//注册鉴权Api
	PrivateGroup := Router.RouterGroup.Group("/api")
	PrivateGroup.Use(middleware.JWTAuth())
	{

	}
	common.GlobalLog.Info("router register success")
	return Router
}
