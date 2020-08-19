package middleware

import (
	gresp "github.com/DowneyL/now/packages/http/response"
	"github.com/DowneyL/now/packages/locales"
	"github.com/DowneyL/now/packages/util"
	"github.com/gin-gonic/gin"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			gresp.FailedError(c, locales.MustTransRespError("token_not_exists"))
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			gresp.FailedError(c, locales.MustTransRespError("token_parse_error"))
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			gresp.FailedError(c, locales.MustTransRespError("token_already_expired"))
			c.Abort()
			return
		}

		c.Next()
	}
}
