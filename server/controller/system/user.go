package system

import (
	"github.com/GanweiYee/ducker/server/entity/system"
	"github.com/GanweiYee/ducker/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (this *UserController) Add(context *gin.Context) {
	var user system.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if err := service.ServiceApp.UserService.Add(&user); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, "successfully")
}
