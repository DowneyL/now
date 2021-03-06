package locales

import (
	"errors"
	"fmt"
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

func GetLanguage() string {
	return Translator.languageTag.String()
}

func (tran *Translate) GetLanguageTag() language.Tag {
	return tran.languageTag
}

func (tran *Translate) MustTrans(messageId string) string {
	return tran.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageId,
	})
}

func (tran *Translate) MustTransAsError(messageId string) error {
	message := tran.MustTrans(messageId)

	return errors.New(message)
}

func MustTransAsError(messageId string) error {
	return Translator.MustTransAsError(messageId)
}

func MustTransRespError(messageId string) error {
	messageId = fmt.Sprintf("response.%s", messageId)
	return MustTransAsError(messageId)
}
