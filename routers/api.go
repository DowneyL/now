package routers

import (
	v1 "github.com/DowneyL/now/controllers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiV1(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/", func(c *gin.Context) {
			c.Data(http.StatusOK, "text/plain", []byte("Hello gin"))
		})
		apiV1.POST("/register", v1.Register)
	}
}
