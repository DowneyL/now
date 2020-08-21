package uv

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	"log"
)

// Register english translation
func RegisterEnTranslation(validate *validator.Validate) {
	locale := en.New()
	uni := ut.New(locale, locale)
	Trans, _ = uni.GetTranslator("en")
	if err := enTranslation.RegisterDefaultTranslations(validate, Trans); err != nil {
		log.Fatalf("universal Trans register failed: %v\n", err)
	}
	if err := registerLocaleTranslation(validate, LanguageTag, "{0} is a wrong language type"); err != nil {
		log.Fatalln("register locale translation failed")
	}
}
