package dao

import (
	"bankaccountapi/model"
	"fmt"

	"github.com/globalsign/mgo/bson"
)

const (
	COLLECTION_Bankaccount = "bankaccount"
)

func (u *UsersDAO) FindBankAccountAll() ([]model.BankAccout, error) {
	var bankaccounts []model.BankAccout
	err := db.C(COLLECTION_Bankaccount).Find(bson.M{}).All(&bankaccounts)
	fmt.Printf("%#v\n", bankaccounts)
	return bankaccounts, err
}

func (u *UsersDAO) InsertBankAccount(bankaccount model.BankAccout) error {
	err := db.C(COLLECTION_Bankaccount).Insert(&bankaccount)
	fmt.Printf("%#v\n", bankaccount)
	return err
}

func (u *UsersDAO) DeleteBankAccount(bankaccount model.BankAccout) error {
	err := db.C(COLLECTION_Bankaccount).Remove(&bankaccount)
	fmt.Printf("%#v\n", bankaccount)
	return err
}
