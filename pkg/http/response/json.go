package gresp

import (
	"github.com/DowneyL/now/pkg/uv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type response struct {
	Code    Code        `json:"code"`
	Message string      `json:"message,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"module"`
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

func bindErrorResponse(err error) *response {
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

func Error(c *gin.Context, code Code, message string, messages []string) {
	c.JSON(http.StatusInternalServerError, errorResponse(code, message, messages))
}

func ServerError(c *gin.Context, err error) {
	Error(c, Internal, err.Error(), nil)
}

func FailedError(c *gin.Context, err error) {
	Error(c, Failed, err.Error(), nil)
}

func BindError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, bindErrorResponse(err))
}

func InvalidArgumentError(c *gin.Context, err error) {
	Error(c, InvalidArgument, err.Error(), nil)
}
