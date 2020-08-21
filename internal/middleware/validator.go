package middleware

import (
	"github.com/DowneyL/now/pkg/locales"
	"github.com/DowneyL/now/pkg/uv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
)

// Used by gin middleware
func UniversalValidator(validate *validator.Validate) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch locales.Translator.GetLanguageTag() {
		case language.English, language.AmericanEnglish, language.BritishEnglish:
			uv.RegisterEnTranslation(validate)
		case language.Chinese, language.SimplifiedChinese, language.TraditionalChinese:
			uv.RegisterZhTranslation(validate)
		case language.Japanese:
			uv.RegisterJaTranslation(validate)
		default:
			uv.RegisterZhTranslation(validate)
		}
	}
}
