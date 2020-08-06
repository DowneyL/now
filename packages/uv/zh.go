package uv

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
	"log"
)

// Register english translation
func RegisterZhTranslation() {
	locale := zh.New()
	uni := ut.New(locale, locale)
	Trans, _ = uni.GetTranslator("zh")
	if err := zhTranslation.RegisterDefaultTranslations(validate, Trans); err != nil {
		log.Fatalf("universal Trans register failed: %v\n", err)
	}
}