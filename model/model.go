package model

import "github.com/globalsign/mgo/bson"

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	FirstName   string        `bson:"first_name" json:"first_name"`
	LastName    string        `bson:"last_name" json:"last_name"`
	BankAccount []BankAccout  `bson:"bank_account" json:"bank_account"`
}

type BankAccout struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	UserID        bson.ObjectId `bson:"user_id" json:"user_id"`
	AccountNumber string        `bson:"account_number" json:"account_number"`
	Balance       float64       `bson:"balance" json:"balance"`
}
