package model

import "github.com/globalsign/mgo/bson"

type Movie struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	FirstName string        `bson:"name" json:"first_name"`
	LastName  string        `bson:"cover_image" json:"last_name"`
}
