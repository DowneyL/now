package routers

import (
	"github.com/DowneyL/now/packages/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(configs.Conf.Mode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 静态文件
	r.StaticFS("upload/images", http.Dir(configs.Conf.GetImageUploadPath()))

	// 路由组
	apiV1(r)

	return r
}
