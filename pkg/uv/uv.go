// go-playground/validator with multi-language

package uv

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

var Trans ut.Translator

const (
	LanguageTag string = "language"
)

func translationFunc(tag string) validator.TranslationFunc {
	return func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())

		return t
	}
}

func registerTranslation(tag, text string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		return ut.Add(tag, text, override)
	}
}

func registerLocaleTranslation(validate *validator.Validate, tag string, text string) error {
	return validate.RegisterTranslation(tag, Trans, registerTranslation(tag, text, true), translationFunc(tag))
}

func GetBindingValidate() *validator.Validate {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatalln("binding validator is not correct")
	}

	_ = validate.RegisterValidation(LanguageTag, Language)
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return getValidatorFieldName(field)
	})

	return validate
}

func getValidatorFieldName(field reflect.StructField) string {
	getFieldName := func(str string) string {
		return strings.SplitN(field.Tag.Get(str), ",", 2)[0]
	}
	name := getFieldName("json")
	if name == "" {
		name = getFieldName("form")
	}
	if name == "-" {
		return ""
	}
	return name
}
