package middlewares

import (
	"net/http"

	"github.com/MHafizAF/bookself-api/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValidate(c)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()

			return
		}

		c.Next()
	}
}
