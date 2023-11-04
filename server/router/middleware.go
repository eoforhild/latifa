package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.SplitN(c.GetHeader("Authorization"), " ", 2)

		if len(token) != 2 || token[0] != "Bearer" {
			c.Header("WWW-Authenticate", "Bearer")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "The required authorization heads were not present in the request.",
			})

			return
		}

		// Pass up further along the context.
		c.Set("Authorization", token)

		c.Next()
	}
}
