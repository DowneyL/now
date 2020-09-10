package api

import (
	"github.com/DowneyL/now/internal/models"
	"github.com/DowneyL/now/internal/services"
	gresp "github.com/DowneyL/now/pkg/http/response"
	"github.com/DowneyL/now/pkg/locales"
	"github.com/DowneyL/now/pkg/util"
	"github.com/gin-gonic/gin"
)

var registerInfo models.RegisterInfo

func Register(c *gin.Context) {
	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		gresp.BindError(c, err)
		return
	}

	if user := models.FindUserByName(registerInfo.Name); user.ID != 0 {
		gresp.FailedError(c, locales.MustTransRespError("username_already_exists"))
		return
	}

	user, auth, err := services.Register(registerInfo.Name, registerInfo.Password)
	if err != nil {
		gresp.ServerInternalError(c)
		return
	}

	gresp.Success(c, gin.H{
		"user": user,
		"auth": auth,
	})
}

func Login(c *gin.Context) {
	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		gresp.BindError(c, err)
		return
	}

	user := models.FindUserByName(registerInfo.Name)
	if user.ID == 0 {
		gresp.FailedError(c, locales.MustTransRespError("user_not_exists"))
		return
	}

	if user.State == models.InvalidState {
		gresp.FailedError(c, locales.MustTransRespError("forbidden_user"))
		return
	}

	if flag := util.ConfirmPassword(user.Password, registerInfo.Password); !flag {
		gresp.FailedError(c, locales.MustTransRespError("wrong_password"))
		return
	}

	auth, err := util.GenerateAuth(user.Name, user.Password)
	if err != nil {
		gresp.ServerInternalError(c)
		return
	}

	gresp.Success(c, gin.H{
		"user": user,
		"auth": auth,
	})
}
