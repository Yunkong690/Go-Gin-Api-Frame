package v1

import "personFrame/service/common"

type ServiceGroupV1 struct {
	UserService
	JwtService common.JwtService
}

var ServiceGroupV1App = new(ServiceGroupV1)
