package uv

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
	"log"
)

// Register english translation
func RegisterZhTranslation(validate *validator.Validate) {
	locale := zh.New()
	uni := ut.New(locale, locale)
	Trans, _ = uni.GetTranslator("zh")
	if err := zhTranslation.RegisterDefaultTranslations(validate, Trans); err != nil {
		log.Fatalf("universal Trans register failed: %v\n", err)
	}
	if err := registerLocaleTranslation(validate, LanguageTag, "{0} 不是正确的语言类型"); err != nil {
		log.Fatalln("register locale translation failed")
	}
}
