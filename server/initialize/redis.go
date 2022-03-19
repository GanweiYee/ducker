package initialize

import (
	"context"
	"fmt"
	"github.com/GanweiYee/feuer/server/global"
	"github.com/go-redis/redis/v8"
)

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.Config.Redis.Ip, global.Config.Redis.Port), // use default Addr
		Password: global.Config.Redis.Password,                                           // no password set
		DB:       global.Config.Redis.DB,                                                 // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pong)
		global.RedisDb = client
	}
}
