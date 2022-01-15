package service

import "github.com/GanweiYee/ducker/server/service/system"

type ServiceGroup struct {
	system.ServiceGroup
}

var ServiceApp = new(ServiceGroup)
