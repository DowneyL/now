package middleware

import (
	"github.com/DowneyL/now/packages/configs"
	"github.com/DowneyL/now/packages/locales"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func Local() gin.HandlerFunc {
	config := configs.New()
	return func(c *gin.Context) {
		lang := c.DefaultQuery("lang", config.GetDefaultLanguage())
		languageTag, err := language.Parse(lang)
		if err != nil {
			languageTag = language.MustParse(config.GetDefaultLanguage())
		}

		locales.SetLanguageTag(languageTag)
	}
}