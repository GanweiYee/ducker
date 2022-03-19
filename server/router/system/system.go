package system

import (
	"github.com/GanweiYee/feuer/server/controller"
	"github.com/GanweiYee/feuer/server/utils"
)

var (
	baseController = controller.App.BaseController
	userController = controller.App.UserController
)

func module() string {
	type i struct{}
	return utils.PkgName(i{})
}
