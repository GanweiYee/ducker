package global

import (
	"github.com/GanweiYee/feuer/server/config"
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

var (
	DbEngine *xorm.Engine
	RedisDb  *redis.Client
	Config   *config.Config
)
