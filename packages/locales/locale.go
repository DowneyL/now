package locales

import (
	"github.com/DowneyL/now/packages/configs"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

var LanguageTag language.Tag

func Local() gin.HandlerFunc {
	return func(c *gin.Context) {
		setup(c)
	}
}

func setup(c *gin.Context) {
	config := configs.New()
	lang := c.DefaultQuery("lang", config.Server.Lang)
	tag, err := language.Parse(lang)
	if err != nil {
		tag = language.MustParse(config.Server.Lang)
	}
	LanguageTag = tag
}
