package uv

import (
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	jaTranslation "github.com/go-playground/validator/v10/translations/ja"
	"log"
)

// Register english translation
func RegisterJaTranslation(validate *validator.Validate) {
	locale := ja.New()
	uni := ut.New(locale, locale)
	Trans, _ = uni.GetTranslator("ja")
	if err := jaTranslation.RegisterDefaultTranslations(validate, Trans); err != nil {
		log.Fatalf("universal Trans register failed: %v\n", err)
	}
	if err := registerLocaleTranslation(validate, LanguageTag, "{0} is a wrong language type"); err != nil {
		log.Fatalln("register locale translation failed")
	}
}
