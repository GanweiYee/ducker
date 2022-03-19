package controller

import "github.com/GanweiYee/feuer/server/controller/system"

type Group struct {
	system.ControllerGroup
}

var App = new(Group)
