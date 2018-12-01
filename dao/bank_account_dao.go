package dao

import (
	"bankaccountapi/model"
	"fmt"

	"github.com/globalsign/mgo/bson"
)

func (u *UsersDAO) FindBankAccountAll() ([]model.BankAccout, error) {
	var bankaccounts []model.BankAccout
	err := db.C(COLLECTION).Find(bson.M{}).All(&bankaccounts)
	fmt.Printf("%#v\n", bankaccounts)
	return bankaccounts, err
}

func (u *UsersDAO) InsertBankAccount(bankaccount model.BankAccout) error {
	err := db.C(COLLECTION).Insert(&bankaccount)
	fmt.Printf("%#v\n", bankaccount)
	return err
}

func (u *UsersDAO) DeleteBankAccount(bankaccount model.BankAccout) error {
	err := db.C(COLLECTION).Remove(&bankaccount)
	fmt.Printf("%#v\n", bankaccount)
	return err
}
