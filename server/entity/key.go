package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Key struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Ik       string             `bson:"ik"`
	Spk      string             `bson:"spk"`
	SpkSig   string             `bson:"spk_sig"`
	PqSpk    string             `bson:"pq_sqk"`
	PqSpkSig string             `bson:"pq_spk_sig"`
	Opk      [32]string         `bson:"opk"`
	PqOpk    [32]string         `bson:"pq_opk"`
	PqOpkSig [32]string         `bson:"pg_opk_sig"`
}
