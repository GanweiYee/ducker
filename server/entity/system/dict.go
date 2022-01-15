package system

import (
	"github.com/GanweiYee/ducker/server/entity"
)

type Dict struct {
	entity.Base `xorm:"extends"`
}

func (receiver *Dict) TableName() string {
	return entity.TableName(*receiver)
}

func init() {
	entity.RegisterDbTable(new(Dict))
}
