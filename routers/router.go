package routers

import (
	"github.com/DowneyL/now/middlewares"
	"github.com/DowneyL/now/packages/configs"
	"github.com/DowneyL/now/packages/uv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func InitRouter() *gin.Engine {
	config := configs.New()
	bundle := getI18nBundle()
	validate := getBindingValidate()
	gin.SetMode(config.Mode)
	r := gin.New()
	// 注册中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translator(bundle))
	r.Use(middleware.UniversalValidator(validate))
	// 静态文件
	r.StaticFS(config.GetImageUploadPath(), http.Dir(config.GetFullImageUploadPath()))
	// 路由
	apiV1(r)

	return r
}

func getI18nBundle() *i18n.Bundle {
	config := configs.New()
	bundle := i18n.NewBundle(language.English)
	for _, file := range config.GetNeedLoadLangFile() {
		bundle.MustLoadMessageFile(file)
	}

	return bundle
}

func getBindingValidate() *validator.Validate {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatalln("binding validator is not correct")
	}

	_ = validate.RegisterValidation(uv.NotExistsTag, uv.NotExists)
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return getValidatorFieldName(field)
	})

	return validate
}

func getValidatorFieldName(field reflect.StructField) string {
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
}
