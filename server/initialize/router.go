package initialize

import (
	_ "github.com/GanweiYee/feuer/server/docs"
	"github.com/GanweiYee/feuer/server/router"
	// 路由引入
	_ "github.com/GanweiYee/feuer/server/router/system"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
	routerEngine := gin.Default()

	routerV := routerEngine.Group("v1")

	for module, routes := range router.RouteGroup {
		moduleRouter := routerV.Group(module)
		for _, route := range routes {
			route(moduleRouter)
		}
	}

	swagger(routerEngine)

	return routerEngine
}

//对接swagger
func swagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
