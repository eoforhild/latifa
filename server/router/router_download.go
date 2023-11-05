package router

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func getFileDownload(c *gin.Context) {
	file := c.Param("file")

	mongodb := ExtractMongoClient(c)

	db := mongodb.Database("latifa_info")
	fs, err := gridfs.NewBucket(
		db,
	)

	// Open the file for reading from MongoDB
	downloadStream, err := fs.OpenDownloadStreamByName(file)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer downloadStream.Close()

	// Read the file data into a byte slice
	data, err := ioutil.ReadAll(downloadStream)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the response headers for the file download
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file))
	c.Data(http.StatusOK, "application/octet-stream", data)

}
