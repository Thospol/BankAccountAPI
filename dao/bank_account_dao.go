package dao

import (
	"bankaccountapi/model"
	"fmt"
)

func (u *UsersDAO) InsertBankOfAcount(user model.User) error {
	//user.BankAccount
	err := db.C(COLLECTION).Insert(&user)
	fmt.Printf("%#v\n", user)
	return err
}
