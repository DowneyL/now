package uv

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	"log"
)

// Register english translation
func RegisterEnTranslation() {
	locale := en.New()
	uni := ut.New(locale, locale)
	Trans, _ = uni.GetTranslator("en")
	if err := enTranslation.RegisterDefaultTranslations(validate, Trans); err != nil {
		log.Fatalf("universal Trans register failed: %v\n", err)
	}
}
