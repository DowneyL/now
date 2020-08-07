package gresp

import (
	"github.com/DowneyL/now/packages/locales"
	uv "github.com/DowneyL/now/packages/uv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data"`
}

func defaultResp(data interface{}) *response {
	return &response{
		Code:    0,
		Message: "",
		Errors:  nil,
		Data:    data,
	}
}

func errorResp(code int, message string, errors []string) *response {
	return &response{
		Code:    code,
		Message: message,
		Errors:  errors,
		Data:    gin.H{},
	}
}

func argumentError(err error) *response {
	errors := err.(validator.ValidationErrors)
	errorStrList := make([]string, len(errors))
	for i := 0; i < len(errors); i++ {
		errorStrList[i] = errors[i].Translate(uv.Trans)
	}

	message := locales.Translator.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "hello",
	})

	return errorResp(-1, message, errorStrList)
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, defaultResp(data))
}

func InvalidArgumentError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, argumentError(err))
}
