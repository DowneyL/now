package middleware

import (
	"github.com/DowneyL/now/models"
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

		user := models.FindUserByName(claims.Username)
		if user.State == models.InvalidState {
			gresp.FailedError(c, locales.MustTransRespError("forbidden_user"))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
