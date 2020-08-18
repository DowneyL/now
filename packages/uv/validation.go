package uv

import (
	"fmt"
	"github.com/DowneyL/now/models"
	"github.com/go-playground/validator/v10"
)

func NotExists(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	topType := fl.Top().Type().String()
	fieldName := fl.FieldName()
	fmt.Println(topType, fieldName)
	switch {
	case topType == "*v1.UserJson" && fieldName == "name":
		return models.FindUserByName(value).ID == 0
	default:
		return true
	}
}
