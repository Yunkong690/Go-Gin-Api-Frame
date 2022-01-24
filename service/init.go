package service

import v1 "personFrame/service/v1"

type ServiceGruop struct {
	ServiceGroupV1 v1.ServiceGroupV1
}

var ServiceGruopApp = new(ServiceGruop)
