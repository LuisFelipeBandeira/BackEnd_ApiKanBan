package middlewares

import (
	"net/http"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/login" {
			return
		}

		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := header[len(Bearer_schema):]

		if !services.ValidateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
