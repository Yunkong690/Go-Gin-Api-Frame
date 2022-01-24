package v1

import (
	"github.com/gin-gonic/gin"
	"personFrame/api"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := api.ApiGroupApp.ApiGroupV1.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
	}
	return baseRouter
}
