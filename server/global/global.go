package global

import (
	"github.com/GanweiYee/ducker/server/config"
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

var (
	DbEngine *xorm.Engine
	RedisDb  *redis.Client
	Config   *config.Config
)
