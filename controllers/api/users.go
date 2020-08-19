package api

import (
	"github.com/DowneyL/now/models"
	gresp "github.com/DowneyL/now/packages/http/response"
	"github.com/DowneyL/now/packages/locales"
	"github.com/DowneyL/now/packages/util"
	"github.com/DowneyL/now/services"
	"github.com/gin-gonic/gin"
)

var json UserJSON

type UserJSON struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	if err := c.ShouldBindJSON(&json); err != nil {
		gresp.BindError(c, err)
		return
	}

	if user := models.FindUserByName(json.Name); user.ID != 0 {
		gresp.FailedError(c, locales.MustTransRespError("username_already_exists"))
		return
	}

	user, auth, err := services.Register(json.Name, json.Password)
	if err != nil {
		gresp.ServerError(c, locales.MustTransRespError("internal"))
		return
	}

	gresp.Success(c, gin.H{
		"user": user,
		"auth": auth,
	})
}

func Login(c *gin.Context) {
	if err := c.ShouldBindJSON(&json); err != nil {
		gresp.BindError(c, err)
		return
	}

	user := models.FindUserByName(json.Name)
	if user.ID == 0 {
		gresp.FailedError(c, locales.MustTransRespError("user_not_exists"))
		return
	}

	if flag := util.ConfirmPassword(user.Password, json.Password); !flag {
		gresp.FailedError(c, locales.MustTransRespError("wrong_password"))
		return
	}

	auth, err := util.GenerateAuth(user.Name, user.Password)
	if err != nil {
		gresp.ServerError(c, locales.MustTransRespError("internal"))
		return
	}

	gresp.Success(c, gin.H{
		"user": user,
		"auth": auth,
	})
}
