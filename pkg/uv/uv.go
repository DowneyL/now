// go-playground/validator with multi-language

package uv

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Trans ut.Translator

const (
	NotExistsTag string = "not_exists"
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
