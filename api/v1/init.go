package v1

import "personFrame/service"

// ApiGroupV1 注册v1版本的路由组
type ApiGroupV1 struct {
	BaseApi
}

var (
	userService = service.ServiceGruopApp.ServiceGroupV1.UserService
	jwtService  = service.ServiceGruopApp.ServiceGroupV1.JwtService
)
