package uv

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
)

func Language(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	tag := language.Make(s)
	return tag.String() != "und"
}
