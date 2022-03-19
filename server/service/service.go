package service

import "github.com/GanweiYee/feuer/server/service/system"

type Group struct {
	system.ServiceGroup
}

var App = new(Group)
