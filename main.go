package main

import (
	"bankaccountapi/dao"
	"bankaccountapi/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

var (
	//config = Config{}
	daos = dao.UsersDAO{}
	r    = mux.NewRouter()
)

func AllUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "AllUserEndPoint !")
}

func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "FindUserEndpoint !")
}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := daos.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UpdateUserEndPoint !")
}

func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DeleteUserEndPoint !")
}

func main() {
	InitRoute()
}
