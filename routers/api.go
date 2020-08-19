package routers

import (
	apiController "github.com/DowneyL/now/controllers/api"
	v1 "github.com/DowneyL/now/controllers/api/v1"
	middleware "github.com/DowneyL/now/middlewares"
	"github.com/gin-gonic/gin"
)

func api(r *gin.Engine) {
	group := r.Group("/api")
	{
		group.POST("/user", apiController.Register)
		group.POST("/user/login", apiController.Login)
	}
}

func apiV1(r *gin.Engine) {
	group := r.Group("/api/v1")
	group.Use(middleware.Jwt())
	{
		group.GET("/database/migrate", v1.Migrate)
	}
}
