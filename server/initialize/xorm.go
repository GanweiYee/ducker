package initialize

import (
	"fmt"
	"github.com/GanweiYee/ducker/server/entity"
	"github.com/GanweiYee/ducker/server/global"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func Xorm() {
	var err error

	var driverName = "mysql"
	var dataSourceName = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		global.Config.Mysql.Username,
		global.Config.Mysql.Password,
		global.Config.Mysql.Ip,
		global.Config.Mysql.Port,
		global.Config.Mysql.Database,
	)

	global.DbEngine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	global.DbEngine.ShowSQL(true) // 日志显示执行sql

	//同步表结构
	if global.Config.Switchs.SyncDb {
		err = global.DbEngine.Sync2(entity.DbTables...)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
	}

}
