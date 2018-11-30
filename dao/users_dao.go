package dao

import (
	"bankaccountapi/model"
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type UsersDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (u *UsersDAO) FindAll() ([]model.User, error) {
	var users []model.User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

func (u *UsersDAO) FindById(id string) (model.User, error) {
	var user model.User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (u *UsersDAO) Insert(user model.User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (u *UsersDAO) Delete(user model.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func (u *UsersDAO) Update(user model.User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
