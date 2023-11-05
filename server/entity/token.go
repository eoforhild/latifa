package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Token  string             `bson:"token"`
}
