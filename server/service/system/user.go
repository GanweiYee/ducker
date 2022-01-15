package system

import (
	systementity "github.com/GanweiYee/ducker/server/entity/system"
	"github.com/GanweiYee/ducker/server/global"
	"github.com/GanweiYee/ducker/server/utils"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
}

func (receiver *UserService) Add(user *systementity.User) (err error) {
	user.ID = uuid.NewV4().String()
	user.Password, _ = utils.PasswordNew(user.Password)
	_, err = global.DbEngine.InsertOne(user)
	return
}

func (receiver *UserService) Delete(ids ...string) (err error) {
	_, err = global.DbEngine.In("id", ids).Delete(&systementity.User{})
	return
}

func (receiver *UserService) Update(user *systementity.User) (err error) {
	cols := []string{"name", "email", "phone"}
	_, err = global.DbEngine.Cols(cols...).Update(user)
	return
}

func (receiver *UserService) List(params map[string]interface{}) {

}
