package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestError struct {
	err error
}

func NewError(err error) *RequestError {
	return &RequestError{err: err}
}

func (e *RequestError) Abort(c *gin.Context) {
	if errors.Is(e.err, mongo.ErrNoDocuments) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "The requested resource could not be found in our records.",
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": "An unexpected error was encountered while processing this request.",
	})
}
