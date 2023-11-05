package router

import (
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewClient(mongodb mongo.Client) *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // HACK

	router := gin.New()
	router.Use(AttachMongoClient(mongodb))
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		log.WithFields(log.Fields{
			"client_ip":   params.ClientIP,
			"user_agent":  params.Request.UserAgent(),
			"latency":     params.Latency,
			"status_code": params.StatusCode}).Infof("%s %s", params.MethodColor()+params.Method+params.ResetColor(), params.Path)

		return ""
	}))

	router.POST("/login", postLogin)
	router.POST("/logout", RequireAuthorization(), UserExists(), postLogout)
	router.POST("/register", postRegister)

	gated := router.Group("")
	gated.Use(RequireAuthorization(), UserExists())
	{
		gated.POST("/files/upload", postFileUpload)
		gated.GET("/files/:file/download", getFileDownload)
		gated.GET("/files", getFiles)
		gated.DELETE("/files/:file", deleteFile)

		gated.POST("/requests/email/:email", postRequestEmail)
		gated.POST("/requests/:request/approve", postRequestApprove)
		gated.GET("/requests/pending", getRequestPending)
	}

	return router
}
