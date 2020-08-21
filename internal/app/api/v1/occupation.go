package v1

import (
	"github.com/DowneyL/now/internal/models"
	gresp "github.com/DowneyL/now/pkg/http/response"
	"github.com/DowneyL/now/pkg/locales"
	"github.com/gin-gonic/gin"
)

var req occupationRequest

type occupationRequest struct {
	Code    string `json:"code" binding:"required,ascii"`
	Lang    string `json:"lang" binding:"required,language"`
	Content string `json:"content" binding:"required"`
}

func CreateOccupation(c *gin.Context) {
	if err := c.ShouldBind(&req); err != nil {
		gresp.BindError(c, err)
		return
	}

	occupation, err := models.CreateOccupation(req.Code, req.Lang, req.Content)
	if err != nil {
		gresp.FailedError(c, locales.MustTransRespError("failed"))
		return
	}

	gresp.Success(c, occupation)
}
