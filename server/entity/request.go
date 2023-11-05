package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	FromID   primitive.ObjectID `json:"from_id" bson:"from_id"`
	ToID     primitive.ObjectID `json:"to_id" bson:"to_id"`
	Accepted bool               `json:"accepted" bson:"accepted"`
	Pending  bool               `json:"pending" bson:"pending"`
}
