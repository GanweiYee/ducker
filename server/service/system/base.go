package system

import (
	systementity "github.com/GanweiYee/ducker/server/entity/system"
	"github.com/GanweiYee/ducker/server/global"
	requestmodel "github.com/GanweiYee/ducker/server/model/request"
	"github.com/GanweiYee/ducker/server/utils"
	"github.com/pkg/errors"
)

type BaseService struct {
}

func (receiver *BaseService) Login(loginForm requestmodel.LoginForm) (user *systementity.User, err error) {
	user = &systementity.User{
		Account: loginForm.Account,
	}

	_, err = global.DbEngine.Get(user)
	if err != nil {
		return
	}

	if !utils.PasswordValidate(user.Password, loginForm.Password) {
		return nil, errors.New("password is failure")
	}

	return
}
