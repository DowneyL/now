package v1

import (
	"fmt"
	"github.com/DowneyL/now/packages/universal-validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func Register(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		errs := err.(validator.ValidationErrors)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, map[string]string{
			"err": fmt.Sprintf("%v", errs.Translate(uv.Trans)),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{})
}
