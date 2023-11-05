package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	FromID   primitive.ObjectID `json:"from_id" bson:"from_id"`
	ToID     primitive.ObjectID `json:"to_id" bson:"to_id"`
	FileUuid uuid.UUID          `bson:"file_uuid"`
}
