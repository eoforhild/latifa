package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Uuid uuid.UUID          `bson:"uuid"`
}
