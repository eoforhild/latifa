package router

import (
	"io"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func postFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadPath := "uploads/"
	if err := c.SaveUploadedFile(file, uploadPath+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mongodb := ExtractMongoClient(c)

	// Open the file for reading
	filePath := uploadPath + file.Filename
	fileData, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer fileData.Close()

	db := mongodb.Database("latifa_info")
	fs, err := gridfs.NewBucket(
		db, // You can customize options here
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Open the file for reading and insert it into MongoDB
	uploadStream, err := fs.OpenUploadStream(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer uploadStream.Close()

	_, err = io.Copy(uploadStream, fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
