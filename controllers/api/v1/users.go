package v1

import (
	"github.com/DowneyL/now/models"
	gresp "github.com/DowneyL/now/packages/http/response"
	"github.com/DowneyL/now/packages/locales"
	"github.com/gin-gonic/gin"
)

type UserJSON struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var json UserJSON
	var translator = locales.Translator
	if err := c.ShouldBindJSON(&json); err != nil {
		gresp.BindError(c, err)
		return
	}

	if user := models.FindUserByName(json.Name); user.ID != 0 {
		gresp.FailedError(c, translator.MustTransAsError("response.username_already_exists"))
		return
	}

	user := models.CreateUser(json.Name, json.Password)
	if user.ID == 0 {
		gresp.FailedError(c, translator.MustTransAsError("response.failed"))
		return
	}
	gresp.Success(c, user)
}

func Login(c *gin.Context) {

}
