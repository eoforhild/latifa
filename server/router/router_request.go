package router

import (
	"context"
	"latifa/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postRequestEmail(c *gin.Context) {
	emailParam := c.Param("email")

	fromUser := ExtractUser(c)

	filter := bson.M{
		"email": emailParam,
	}

	var toUser entity.User
	mongodb := ExtractMongoClient(c)
	database := mongodb.Database("latifa_info")
	usercoll := database.Collection("users")
	err := usercoll.FindOne(context.Background(), filter).Decode(&toUser)
	if err != nil {
		NewError(err).Abort(c)
	}

	newRequest := bson.M{
		"handshakes":    nil,
		"from_username": fromUser.Username,
		"to_id":         toUser.ID,
		"to_username":   toUser.Username,
		"accepted":      false,
		"pending":       true,
	}

	requestcoll := database.Collection("requests")
	_, err = requestcoll.InsertOne(context.TODO(), newRequest)

	if err != nil {
		NewError(err).Abort(c)
	}

	c.Status(http.StatusNoContent)
}

func postRequestApprove(c *gin.Context) {
	user := ExtractUser(c)

	requestIdParam := c.Param("request")

	requestId, err := primitive.ObjectIDFromHex(requestIdParam)
	if err != nil {
		NewError(err).Abort(c)
	}

	filter := bson.M{
		"_id":   requestId,
		"to_id": user.ID,
	}

	print(requestIdParam)

	update := bson.M{
		"$set": bson.M{
			"accepted": true,
			"pending":  false,
		},
	}

	mongodb := ExtractMongoClient(c)
	database := mongodb.Database("latifa_info")
	requestcoll := database.Collection("requests")

	var updatedRequest bson.M
	err = requestcoll.FindOneAndUpdate(context.TODO(), filter, update).Decode(&updatedRequest)
	if err != nil {
		NewError(err).Abort(c)
	}

	c.Status(http.StatusNoContent)
}

func getRequestPending(c *gin.Context) {
	user := ExtractUser(c)

	filter := bson.M{
		"to_id":   user.ID,
		"pending": true,
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

func getRequestSentHistory(c *gin.Context) {
	user := ExtractUser(c)

	filter := bson.M{
		"from_id": user.ID,
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

func getRequestReceivedHistory(c *gin.Context) {
	user := ExtractUser(c)

	filter := bson.M{
		"to_id": user.ID,
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

func getRequestApproved(c *gin.Context) {
	user := ExtractUser(c)

	filter := bson.M{
		"to_id":    user.ID,
		"approved": true,
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
