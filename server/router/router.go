package router

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type route func(group *gin.RouterGroup)

var (
	RouteGroup = make(map[string][]route)
	routeMutex sync.Mutex
)

func RegisterRoute(module string, r ...route) {
	routeMutex.Lock()
	defer routeMutex.Unlock()

	if RouteGroup[module] == nil {
		RouteGroup[module] = make([]route, 0)
	}

	RouteGroup[module] = append(RouteGroup[module], r...)
}
