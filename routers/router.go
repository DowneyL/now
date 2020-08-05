package routers

import (
	"github.com/DowneyL/now/packages/configs"
	"github.com/DowneyL/now/packages/gin/locales"
	uv "github.com/DowneyL/now/packages/gin/universal-validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	config := configs.New()
	gin.SetMode(config.Mode)
	r := gin.New()
	// 注册中间件
	r.Use(middlewares()...)
	// 静态文件
	r.StaticFS(config.GetImageUploadPath(), http.Dir(config.GetFullImageUploadPath()))
	// 路由
	apiV1(r)

	return r
}

func middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),
		gin.Recovery(),
		locales.Local(),
		uv.Universal(),
	}
}
