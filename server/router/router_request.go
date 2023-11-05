package router

import (
	"context"
	"latifa/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmailRequest struct {
	Email string `json:"email" binding:"required"`
}

func postRequestEmail(c *gin.Context) {
	var e EmailRequest
	if err := c.BindJSON(&e); err != nil {
		return
	}

	fromUser := ExtractUser(c)

	filter := bson.M{
		"email": e.Email,
	}

	var toUser entity.User
	mongodb := ExtractMongoClient(c)
	database := mongodb.Database("latifa_info")
	usercoll := database.Collection("users")
	err := usercoll.FindOne(context.Background(), filter).Decode(&toUser)
	if err != nil {
		NewError(err).Abort(c)
	}

	request := &entity.Request{
		FromID:   fromUser.ID,
		ToID:     toUser.ID,
		Accepted: false,
		Pending:  true,
	}

	requestcoll := database.Collection("requests")
	_, err = requestcoll.InsertOne(context.TODO(), request)

	if err != nil {
		NewError(err).Abort(c)
	}

	c.Status(http.StatusNoContent)
}

func postRequestApprove(c *gin.Context) {
	user := ExtractUser(c)

	requestIdParam := c.Param("request")

	filter := bson.M{
		"_id":  requestIdParam,
		"ToID": user.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"accepted": true,
			"pending":  false,
		},
	}

	mongodb := ExtractMongoClient(c)
	database := mongodb.Database("latifa_info")
	requestcoll := database.Collection("requests")

	options := options.FindOneAndUpdate().SetUpsert(false)

	var updatedRequest bson.M
	err := requestcoll.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&updatedRequest)
	if err != nil {
		NewError(err).Abort(c)
	}

	c.Status(http.StatusNoContent)
}

func getRequestPending(c *gin.Context) {
	user := ExtractUser(c)

	filter := bson.M{
		"ToID": user.ID,
	}

	mongodb := ExtractMongoClient(c)
	database := mongodb.Database("latifa_info")
	requestcoll := database.Collection("requests")
	cursor, err := requestcoll.Find(context.Background(), filter)
	if err != nil {
		NewError(err).Abort(c)
	}
	defer cursor.Close(context.Background())

	var requests []entity.Request
	for cursor.Next(context.Background()) {
		var request entity.Request
		if err := cursor.Decode(&request); err != nil {
			NewError(err).Abort(c)
		}

		// found document
		requests = append(requests, request)
	}

	if err := cursor.Err(); err != nil {
		NewError(err).Abort(c)
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})
}
