package controller

import "github.com/GanweiYee/ducker/server/controller/system"

type ControllerGroup struct {
	system.ControllerGroup
}

var ControllerApp = new(ControllerGroup)
