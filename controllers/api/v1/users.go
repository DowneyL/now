package v1

import (
	"github.com/gin-gonic/gin"
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
		log.Println(err)
		c.JSON(http.StatusInternalServerError, map[string]string{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{})
}
