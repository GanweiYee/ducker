package system

import (
	"github.com/GanweiYee/ducker/server/router"
	"github.com/gin-gonic/gin"
)

func init() {
	router.RegisterRoute(module(), baseRoute)
}

func baseRoute(r *gin.RouterGroup) {
	baseRouter := r.Group("base")

	baseRouter.POST("/login", baseController.Login)
	baseRouter.POST("/logout", baseController.Logout)
}
