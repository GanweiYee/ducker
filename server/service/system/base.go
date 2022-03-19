package system

import (
	systemEntity "github.com/GanweiYee/feuer/server/entity/system"
	"github.com/GanweiYee/feuer/server/global"
	"github.com/GanweiYee/feuer/server/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type BaseService struct {
}

func (s *BaseService) Login(user *systemEntity.User, password string) (err error) {

	_, err = global.DbEngine.Get(user)
	if err != nil {
		return
	}

	if !utils.PasswordValidate(user.Password, password) {
		return errors.New("password is failure")
	}

	return
}

func (s *UserService) Signup(user *systemEntity.User) (err error) {
	user.ID = uuid.NewV4().String()
	user.Password, _ = utils.PasswordNew(user.Password)

	_, err = global.DbEngine.InsertOne(user)
	if err != nil {
		return
	}

	return
}
