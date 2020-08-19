package v1

import (
	"github.com/DowneyL/now/models"
	gresp "github.com/DowneyL/now/packages/http/response"
	"github.com/DowneyL/now/packages/locales"
	"github.com/gin-gonic/gin"
)

var json AddressJSON

type AddressJSON struct {
	Email string `json:"email" binding:"required,email"`
}

func SetEmail(c *gin.Context) {
	if err := c.ShouldBind(&json); err != nil {
		gresp.BindError(c, err)
		return
	}

	user := c.MustGet("user").(*models.User)
	if user.Email != "" {
		gresp.FailedError(c, locales.MustTransRespError("email_exists"))
		return
	}
	if err := models.UpdateUserEmail(user, json.Email); err != nil {
		gresp.ServerError(c, locales.MustTransRespError("internal"))
		return
	}

	gresp.Success(c, gin.H{})
}
