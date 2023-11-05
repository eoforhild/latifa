package router

import (
	"context"
	"latifa/entity"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
		c.Set("Authorization", token[1])

		c.Next()
	}
}

func UserExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u *entity.User
		var token *entity.Token
		auth_token, ok := c.Get("Authorization")
		if !ok {
			panic("router/middleware: expected authorization heads, not found in request")
		}

		mongodb := ExtractMongoClient(c)
		database := mongodb.Database("latifa_info")
		tokencoll := database.Collection("auth_tokens")

		filter := bson.M{
			"token": auth_token,
		}
		err := tokencoll.FindOne(context.TODO(), filter).Decode(&token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "You are not authorized to access this endpoint.",
			})
			return
		}

		ufilter := bson.M{
			"_id": token.UserId,
		}

		usercoll := database.Collection("users")
		err = usercoll.FindOne(context.TODO(), ufilter).Decode(&u)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "You are not authorized to access this endpoint.",
			})
			return
		}

		c.Set("user", u)
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
	panic("router/middleware: cannot extract mongo client: not present in request context")
}

func ExtractUser(c *gin.Context) *entity.User {
	v, ok := c.Get("user")
	if !ok {
		panic("router/middleware: cannot extract user: not present in request context")
	}
	return v.(*entity.User)
}
