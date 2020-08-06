package middleware

import (
	"github.com/DowneyL/now/packages/locales"
	uv "github.com/DowneyL/now/packages/uv"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

// Used by gin middleware
func UniversalValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := locales.GetLanguageTag()
		switch lang {
		case language.English, language.AmericanEnglish, language.BritishEnglish:
			uv.RegisterEnTranslation()
		case language.Chinese, language.SimplifiedChinese, language.TraditionalChinese:
			uv.RegisterZhTranslation()
		case language.Japanese:
			uv.RegisterJaTranslation()
		default:
			uv.RegisterZhTranslation()
		}
	}
}