package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	UserId primitive.ObjectID `bson:"user_id"`
	Token  string             `bson:"token"`
}
