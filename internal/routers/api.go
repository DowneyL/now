package routers

import (
	v0 "github.com/DowneyL/now/internal/app/api"
	v1 "github.com/DowneyL/now/internal/app/api/v1"
	"github.com/DowneyL/now/internal/middleware"
	"github.com/gin-gonic/gin"
)

func api(r *gin.Engine) {
	group := r.Group("/api")
	{
		group.GET("/database/migrate", v0.Migrate)
		group.POST("/user", v0.Register)
		group.POST("/user/login", v0.Login)
	}
}

func apiV1(r *gin.Engine) {
	group := r.Group("/api/v1")
	group.Use(middleware.Jwt())
	{
		group.PATCH("/user/email", v1.SetEmail)
		group.PATCH("/user", v1.EditUser)
		group.POST("/occupation", v1.CreateOccupation)
	}
}
