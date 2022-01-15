package system

import (
	systementity "github.com/GanweiYee/ducker/server/entity/system"
	systemmodel "github.com/GanweiYee/ducker/server/model"
	requestmodel "github.com/GanweiYee/ducker/server/model/request"
	"github.com/GanweiYee/ducker/server/model/response"
	"github.com/GanweiYee/ducker/server/service"
	"github.com/GanweiYee/ducker/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

// @登录
// @Description login
// @Accept  json
// @Produce json
// @Param   username     path    string     true        "username"
// @Param   passwd     path    string     true        "passwd"
// @Success 200 {string} string	"ok"
// @Router /login [post]
func (receiver *BaseController) Login(context *gin.Context) {
	var loginForm requestmodel.LoginForm

	if err := context.ShouldBindJSON(&loginForm); err != nil {
		context.JSON(http.StatusBadRequest, "请求参数错误")
		return
	}

	user, err := service.ServiceApp.Login(loginForm)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Data{
			Msg: err.Error(),
		})
		return
	}

	receiver.token(context, user)
}
func (receiver *BaseController) Logout(context *gin.Context) {
}

func (receiver *BaseController) token(context *gin.Context, user *systementity.User) {
	j := utils.NewJWT()

	claims := j.CreateClaims(systemmodel.BaseClaims{
		ID:       user.ID,
		Username: user.Nickname,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, response.LoginResponse{
		Token: token,
	})
}
