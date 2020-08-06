// go-playground/validator with multi-language

package uv

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

var (
	Trans    ut.Translator
	validate *validator.Validate
	ok       bool
)

func init() {
	validate, ok = binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatalln("binding validator is not correct")
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		getFieldName := func(str string) string {
			return strings.SplitN(field.Tag.Get(str), ",", 2)[0]
		}
		name := getFieldName("json")
		if name == "" {
			name = getFieldName("form")
		}
		if name == "-" {
			return ""
		}
		return name
	})
}
