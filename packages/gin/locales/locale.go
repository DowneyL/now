package locales

import (
	"github.com/DowneyL/now/packages/configs"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

var (
	config = configs.New()
	languageTag language.Tag
	err         error
)

func LanguageTag() language.Tag {
	if languageTag.String() != "und" {
		return languageTag
	}

	languageTag = language.MustParse(config.Server.Lang)
	return languageTag
}

func Local() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.DefaultQuery("lang", config.Server.Lang)
		languageTag, err = language.Parse(lang)
		if err != nil {
			languageTag = language.MustParse(config.Server.Lang)
		}
	}
}
