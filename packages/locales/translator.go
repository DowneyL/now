package locales

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Translator *Translate

type Translate struct {
	languageTag language.Tag
	*i18n.Localizer
}

func New(bundle *i18n.Bundle, lang string) *Translate {
	Translator = &Translate{
		languageTag: language.MustParse(lang),
		Localizer:   i18n.NewLocalizer(bundle, lang),
	}

	return Translator
}

func (tran *Translate) GetLanguageTag() language.Tag {
	return tran.languageTag
}
