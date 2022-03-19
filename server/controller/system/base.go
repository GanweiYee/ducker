package system

import (
	systemEntity "github.com/GanweiYee/feuer/server/entity/system"
	systemModel "github.com/GanweiYee/feuer/server/model"
	requestModel "github.com/GanweiYee/feuer/server/model/request"
	"github.com/GanweiYee/feuer/server/model/response"
	"github.com/GanweiYee/feuer/server/service"
	"github.com/GanweiYee/feuer/server/utils"
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
func (c *BaseController) Login(context *gin.Context) {
	var reqData requestModel.Login

	if err := context.ShouldBindJSON(&reqData); err != nil {
		context.JSON(http.StatusBadRequest, "请求参数错误")
		return
	}

	user := &systemEntity.User{
		Account: reqData.Account,
	}

	err := service.App.Login(user, reqData.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Data{
			Msg: err.Error(),
		})
		return
	}

	c.token(context, user)
}

func (c *BaseController) Signup(context *gin.Context) {
	var reqData requestModel.Signup
	if err := context.ShouldBindJSON(&reqData); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	user := systemEntity.User{
		Account:  reqData.Account,
		Phone:    reqData.Phone,
		Password: reqData.Password,
	}

	if err := service.App.Signup(&user); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, "successfully")
}

func (c *BaseController) Logout(context *gin.Context) {

}

func (c *BaseController) token(context *gin.Context, user *systemEntity.User) {
	j := utils.NewJWT()

	claims := j.CreateClaims(systemModel.BaseClaims{
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
