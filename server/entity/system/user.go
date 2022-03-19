package system

import (
	"github.com/GanweiYee/feuer/server/entity"
)

func init() {
	entity.RegisterDbTable(new(User))
}

type User struct {
	Account      string `json:"account" xorm:"varchar(25) notnull unique 'account' default('') comment('账号')"`
	Password     string `json:"password" xorm:"varchar(50) notnull 'password' default('') comment('账号')"`
	Nickname     string `json:"nickname" xorm:"varchar(25) unique 'nickname' default('') comment('用户名')"`
	ProfilePhoto string `json:"profile_photo" xorm:"varchar(100) null 'profile_photo' default('') comment('头像')"`
	Phone        string `json:"phone" xorm:"varchar(17) null unique 'phone' default('') comment('手机号码')"`
	Email        string `json:"email" xorm:"varchar(50) null unique 'email' default('') comment('邮箱')"`
	Gender       int    `json:"gender" xorm:"tinyint null 'gender' default(0) comment('性别')"`
	Region       string `json:"region" xorm:"varchar(17) null 'region' default('') comment('所在城市')"`
	entity.Base  `xorm:"extends"`
}

func (receiver *User) TableName() string {
	return entity.TableName(*receiver)
}
