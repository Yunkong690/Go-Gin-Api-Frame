package api

import v1 "personFrame/api/v1"

// ApiGroup  注册Api组
type ApiGroup struct {
	ApiGroupV1 v1.ApiGroupV1
}

var ApiGroupApp = new(ApiGroup)
