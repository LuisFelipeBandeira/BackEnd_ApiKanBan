package middlewares

import (
	"net/http"
	"strings"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/login" {
			return
		}

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := strings.Split(header, " ")[1]

		if err := services.ValidateToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
