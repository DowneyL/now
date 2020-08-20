package middleware

import (
	"github.com/DowneyL/now/pkg/configs"
	"github.com/DowneyL/now/pkg/locales"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Translator(bundle *i18n.Bundle) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Query("lang")
		if lang == "" {
			lang = c.GetHeader("Accept-Language")
		}
		if lang == "" {
			config := configs.New()
			lang = config.GetDefaultLanguage()
		}
		locales.New(bundle, lang)
	}
}
