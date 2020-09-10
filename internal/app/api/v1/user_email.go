package v1

import (
	"github.com/DowneyL/now/internal/models"
	gresp "github.com/DowneyL/now/pkg/http/response"
	"github.com/DowneyL/now/pkg/locales"
	"github.com/gin-gonic/gin"
)

var json AddressJSON

type AddressJSON struct {
	Email string `req:"email" binding:"required,email"`
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
		gresp.ServerInternalError(c)
		return
	}

	gresp.Success(c, gin.H{})
}
