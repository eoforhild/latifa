package router

import (
	"context"
	"latifa/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HandshakeRequest struct {
	RequestId string `json:"request_id" binding:"required"`
}

func getHandshake(c *gin.Context) {
	var h HandshakeRequest
	if err := c.BindJSON(&h); err != nil {
		return
	}

	requestUser := ExtractUser(c)
	mongodb := ExtractMongoClient(c)

	requestId, _ := primitive.ObjectIDFromHex(h.RequestId)

	// find the sender user ID
	filter := bson.M{
		"_id":      requestId, // todo convert to hex
		"to_id":    requestUser.ID,
		"approved": true,
		"pending":  false,
	}

	var request entity.Request

	database := mongodb.Database("latifa_info")
	requestcoll := database.Collection("requests")
	err := requestcoll.FindOne(context.Background(), filter).Decode(&request)
	if err != nil {
		NewError(err).Abort(c)
	}

	var key entity.Key

	keyFilter := bson.M{
		"_id": request.FromID.Hex(),
	}

	keycoll := database.Collection("public_keys")
	err = keycoll.FindOne(context.Background(), keyFilter).Decode(&key)
	if err != nil {
		NewError(err).Abort(c)
	}

	ind := 0
	for i := 31; i >= 0; i-- {
		if key.PqOpk[i] != "0000000000000000000000000000000000000000000000000000000000000000" {
			ind = i
			break
		}
	}

	var PQPK string
	var PQPK_Sig string
	// We have no more pqopk keys
	if ind == -1 {
		PQPK = key.PqSpk
		PQPK_Sig = key.PqSpkSig
	} else {
		PQPK = key.PqOpk[ind]
		PQPK_Sig = key.PqOpkSig[ind]
		key.PqOpk[ind] = "0000000000000000000000000000000000000000000000000000000000000000"
		key.PqOpkSig[ind] = "0000000000000000000000000000000000000000000000000000000000000000"

		update := bson.M{
			"$set": bson.M{
				"pqopk_arr":     key.PqOpk,
				"pqopk_sig_arr": key.PqOpkSig,
			},
		}

		var updatedRequest bson.M
		err = keycoll.FindOneAndUpdate(context.Background(), keyFilter, update).Decode(&updatedRequest)
		if err != nil {
			NewError(err).Abort(c)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"ik":       key.Ik,
		"spk":      key.Spk,
		"spk_sig":  key.SpkSig,
		"pqpk":     PQPK,
		"pqpk_sig": PQPK_Sig})
}

type NewHandshakeRequest struct {
	RequestId string `json:"request_id" binding:"required"`
	Handshake string `json:"handshake" binding:"required"`
}

func postHandshake(c *gin.Context) {
	//var h NewHandshakeRequest

}
