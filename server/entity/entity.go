package entity

import (
	"fmt"
	"github.com/GanweiYee/ducker/server/global"
	"github.com/GanweiYee/ducker/server/utils"
	"reflect"
	"sync"
	"time"
)

type Table interface {
	TableName() string
}

var (
	DbTables     = make([]interface{}, 0)
	dbTablesLock sync.Mutex
)

func RegisterDbTable(dbTable Table) {
	dbTablesLock.Lock()
	defer dbTablesLock.Unlock()

	DbTables = append(DbTables, dbTable)
}

// Base 数据库表基础键
type Base struct {
	ID      string    `json:"id" xorm:"varchar(17) notnull pk 'id' comment('主键id')"`
	Created time.Time `json:"created" xorm:"datetime notnull created 'created' comment('创建时间')"`
	Updated time.Time `json:"updated" xorm:"datetime notnull updated 'updated' comment('更新时间')"`
	Deleted time.Time `json:"deleted" xorm:"datetime deleted 'deleted' comment('是否已删除')"`
	Version int       `json:"version" xorm:"int notnull version 'version' default(1) comment('版本')"`
}

// TableName 统一表名格式
func TableName(receiver any) string {
	return fmt.Sprintf("%s_%s_%s",
		global.Config.App.Name,  //项目名称
		utils.PkgName(receiver), //模块名称
		utils.GonicCasedName(reflect.TypeOf(receiver).Name()),
	)
}
