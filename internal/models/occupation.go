package models

import (
	"github.com/DowneyL/now/pkg/util"
	"golang.org/x/text/language"
)

type Occupation struct {
	BaseModel
	Code    string           `gorm:"unique_index:uk_code_lang;size:50;not null;default:'';comment:'职业CODE'" json:"code"`
	Lang    util.LanguageTag `gorm:"unique_index:uk_code_lang;type:varchar(10);not null;default:0;comment:'语言'" json:"lang"`
	Content string           `gorm:"size:50;not null;default:'';comment:'职业内容'" json:"content"`
}

type UserOccupations []Occupation

func (o *UserOccupations) MarshalJSON() ([]byte, error) {
	contents := make(map[string]string, len(*o))
	for _, v := range *o {
		contents[v.Lang.String()] = v.Content
	}

	return util.ToByteWithJSON(contents)
}

func CreateOccupation(code, lang, content string) (occ *Occupation, err error) {
	occ = &Occupation{
		Code:    code,
		Lang:    util.LanguageTag{Tag: language.Make(lang)},
		Content: content,
	}
	err = WriteDB.Create(occ).Error

	return occ, err
}
