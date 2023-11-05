package entity

import (
	"github.com/google/uuid"
)

type Request struct {
	FileID uuid.UUID `json:"file_id" bson:"file_id"`
	FromID uuid.UUID `json:"from_id" bson:"from_id"`
	ToID   uuid.UUID `json:"to_id" bson:"to_id"`
}
