package uv

import (
	"github.com/go-playground/validator/v10"
)

func NotExists(fl validator.FieldLevel) bool {
	return true
}
