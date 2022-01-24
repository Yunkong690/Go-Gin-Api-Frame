package router

import v1 "personFrame/router/v1"

type RouterGroup struct {
	v1.RouterGroupV1
}

var RouterGroupApp = new(RouterGroup)
