package system

import (
	systemEntity "github.com/GanweiYee/feuer/server/entity/system"
	"github.com/GanweiYee/feuer/server/global"
	"github.com/GanweiYee/feuer/server/utils"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
}

func (s *UserService) Add(user *systemEntity.User) (err error) {
	user.ID = uuid.NewV4().String()
	user.Password, _ = utils.PasswordNew(user.Password)

	_, err = global.DbEngine.InsertOne(user)
	if err != nil {
		return
	}

	return
}

func (s *UserService) Delete(ids ...string) (err error) {
	_, err = global.DbEngine.In("id", ids).Delete(&systemEntity.User{})
	return
}

func (s *UserService) Update(user *systemEntity.User) (err error) {
	cols := []string{"name", "email", "phone"}
	_, err = global.DbEngine.Cols(cols...).Update(user)
	return
}

func (s *UserService) List(params map[string]interface{}) {

}
