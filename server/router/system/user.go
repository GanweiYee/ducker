package system

import (
	"github.com/GanweiYee/feuer/server/router"
	"github.com/gin-gonic/gin"
)

func init() {
	router.RegisterRoute(module(), userRoute)
}

func userRoute(r *gin.RouterGroup) {
	userRouter := r.Group("user")

	userRouter.POST("/register", userController.Register)
}
