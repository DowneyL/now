package gresp

import (
	uv "github.com/DowneyL/now/packages/uv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type response struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data"`
}

func successResponse(data interface{}) *response {
	return &response{
		Code:    OK,
		Message: OK.String(),
		Errors:  nil,
		Data:    data,
	}
}

func errorResponse(code Code, message string, errors []string) *response {
	if message == "" {
		message = code.String()
	}
	if len(errors) == 0 {
		errors = nil
	}
	return &response{
		Code:    code,
		Message: message,
		Errors:  errors,
		Data:    gin.H{},
	}
}

func argumentErrorResponse(err error) *response {
	errors := err.(validator.ValidationErrors)
	errorStrList := make([]string, len(errors))
	for i := 0; i < len(errors); i++ {
		errorStrList[i] = errors[i].Translate(uv.Trans)
	}

	return errorResponse(InvalidArgument, "", errorStrList)
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, successResponse(data))
}

func FailedError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, errorResponse(Failed, err.Error(), nil))
}

func InvalidArgumentError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, argumentErrorResponse(err))
}
