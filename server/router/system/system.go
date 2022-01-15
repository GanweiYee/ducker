package system

import (
	"github.com/GanweiYee/ducker/server/controller"
	"github.com/GanweiYee/ducker/server/utils"
)

var (
	baseController = controller.ControllerApp.BaseController
	userController = controller.ControllerApp.UserController
)

func module() string {
	type i struct{}
	return utils.PkgName(i{})
}
