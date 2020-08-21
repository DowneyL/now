package routers

import (
	"github.com/DowneyL/now/internal/middleware"
	"github.com/DowneyL/now/pkg/configs"
	"github.com/DowneyL/now/pkg/uv"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"net/http"
)

func InitRouter() *gin.Engine {
	config := configs.New()
	bundle := getI18nBundle()
	validate := uv.GetBindingValidate()
	gin.SetMode(config.Mode)
	r := gin.New()
	// 注册中间件
	r.Use(gin.Logger(),
		gin.Recovery(),
		middleware.Translator(bundle),
		middleware.UniversalValidator(validate))
	// 静态文件
	r.StaticFS(config.GetImageUploadPath(), http.Dir(config.GetFullImageUploadPath()))
	// 路由
	api(r)
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
