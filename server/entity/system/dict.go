package system

import (
	"github.com/GanweiYee/feuer/server/entity"
)

func init() {
	entity.RegisterDbTable(new(Dict))
}

type Dict struct {
	entity.Base `xorm:"extends"`
}

func (receiver *Dict) TableName() string {
	return entity.TableName(*receiver)
}
