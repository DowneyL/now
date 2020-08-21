package api

import (
	"github.com/DowneyL/now/internal/models"
	gresp "github.com/DowneyL/now/pkg/http/response"
	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	err := models.Migrate()
	if err != nil {
		gresp.FailedError(c, err)
		return
	}

	gresp.Success(c, gin.H{})
}
