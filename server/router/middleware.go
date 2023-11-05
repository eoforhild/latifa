package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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

func AttachMongoClient(client mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mongo_client", client)
		c.Next()
	}
}

func ExtractMongoClient(c *gin.Context) mongo.Client {
	if v, ok := c.Get("mongo_client"); ok {
		return v.(mongo.Client)
	}
	panic("router/middleware: mongo client not present in context")
}
