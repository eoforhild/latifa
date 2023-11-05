package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Key struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Ik       string             `bson:"ik"`
	Spk      string             `bson:"spk"`
	SpkSig   string             `bson:"spk_sig"`
	PqSpk    string             `bson:"pqsqk"`
	PqSpkSig string             `bson:"pqspk_sig"`
	Opk      [32]string         `bson:"opk_arr"`
	PqOpk    [32]string         `bson:"pqopk_arr"`
	PqOpkSig [32]string         `bson:"pqopk_sig_arr"`
}
