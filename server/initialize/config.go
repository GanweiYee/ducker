package initialize

import (
	"fmt"
	"github.com/GanweiYee/feuer/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Config() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
			return
		}
	})
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
		return
	}
}
