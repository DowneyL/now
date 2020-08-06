package locales

import (
	"github.com/DowneyL/now/packages/configs"
	"golang.org/x/text/language"
)

var languageTag language.Tag

func SetLanguageTag(tag language.Tag) {
	languageTag = tag
}

func GetLanguageTag() language.Tag {
	config := configs.New()
	if languageTag.String() != "und" {
		return languageTag
	}

	languageTag = language.MustParse(config.GetDefaultLanguage())
	return languageTag
}
