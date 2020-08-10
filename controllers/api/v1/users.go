package v1

import (
	gresp "github.com/DowneyL/now/packages/http/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		gresp.InvalidArgumentError(c, err)
		return
	}

	gresp.Success(c, map[string]string{})
}

func Login(c *gin.Context) {
	
}
