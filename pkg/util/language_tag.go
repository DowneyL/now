package util

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/DowneyL/now/pkg/configs"
	"golang.org/x/text/language"
)

type LanguageTag struct {
	language.Tag
}

func (tag *LanguageTag) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, tag.Tag.String())), nil
}

func (tag LanguageTag) UnmarshalJSON(data []byte) error {
	s := string(data)
	tag.Tag = language.Make(s)
	return nil
}

func (tag *LanguageTag) Scan(src interface{}) error {
	if src == nil {
		tag.Tag = language.Make(configs.GetDefaultLanguage())
	}

	switch v := src.(type) {
	case string:
		tag.Tag = language.Make(v)
	case []byte:
		tag.Tag = language.Make(string(v))
	default:
		return errors.New("language convert error")
	}

	return nil
}

func (tag LanguageTag) Value() (driver.Value, error) {
	return tag.String(), nil
}
