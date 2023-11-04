package router

import (
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

func NewClient() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // HACK

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		log.WithFields(log.Fields{
			"client_ip":   params.ClientIP,
			"user_agent":  params.Request.UserAgent(),
			"latency":     params.Latency,
			"status_code": params.StatusCode}).Infof("%s %s", params.MethodColor()+params.Method+params.ResetColor(), params.Path)

		return ""
	}))

	router.POST("/login", postLogin)
	router.POST("/logout", RequireAuthorization(), postLogout)
	router.POST("/register", postRegister)

	router.POST("/files", RequireAuthorization())
	router.GET("/files/:file", RequireAuthorization())
	router.DELETE("/files/:file", RequireAuthorization())

	return router
}
