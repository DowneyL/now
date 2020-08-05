package uv

import (
	"github.com/DowneyL/now/packages/gin/locales"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
	"golang.org/x/text/language"
	"log"
	"reflect"
	"strings"
)

var (
	Trans    ut.Translator
	validate *validator.Validate
	ok       bool
)

// Used by gin middleware
func Universal() gin.HandlerFunc {
	return func(c *gin.Context) {
		initValidate()
		lang := locales.LanguageTag()
		switch lang {
		case language.English, language.AmericanEnglish, language.BritishEnglish:
			enTranslationRegister()
		case language.Chinese, language.SimplifiedChinese, language.TraditionalChinese:
			zhTranslationRegister()
		default:
			zhTranslationRegister()
		}
	}
}

func initValidate() {
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

// Register english translation
func enTranslationRegister() {
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	Trans, _ = uni.GetTranslator("en")
	err := enTranslation.RegisterDefaultTranslations(validate, Trans)
	if err != nil {
		log.Fatalf("universal Trans register failed: %v\n", err)
	}
}

// Register english translation
func zhTranslationRegister() {
	zhLocale := zh.New()
	uni := ut.New(zhLocale, zhLocale)
	Trans, _ = uni.GetTranslator("zh")
	err := zhTranslation.RegisterDefaultTranslations(validate, Trans)
	if err != nil {
		log.Fatalf("universal Trans register failed: %v\n", err)
	}
}
