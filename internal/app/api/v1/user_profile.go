package v1

import (
	"github.com/DowneyL/now/internal/models"
	gresp "github.com/DowneyL/now/pkg/http/response"
	"github.com/gin-gonic/gin"
)

var userProfile models.UserProfile

func EditUser(c *gin.Context) {
	if err := c.ShouldBindJSON(&userProfile); err != nil {
		gresp.BindError(c, err)
		return
	}

	user := c.MustGet("user").(*models.User)
	if err := models.EditUser(user, &userProfile); err != nil {
		gresp.ServerInternalError(c)
		return
	}

	gresp.Success(c, gin.H{})
}
