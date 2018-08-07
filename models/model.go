package models

import "gopkg.in/mgo.v2/bson"

type Cars struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	model string        `bson:"model_name" json:"model_name"`
	year  string        `bson:"year" json:"year"`
	marka string        `bson:"marka" json:"marka"`
}
