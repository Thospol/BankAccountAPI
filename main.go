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
	config = Config{}
	daos   = dao.UsersDAO{}
	r      = mux.NewRouter()
)

func init() {

	config.Read()

	daos.Server = config.Server
	daos.Database = config.Database
	daos.Connect()
	fmt.Println("start Server Port:4525")
}

func main() {
	SetupRoute()
}

func AllUserEndPoint(w http.ResponseWriter, r *http.Request) {

	users, err := daos.FindAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, users)
}

func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	user, err := daos.FindById(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	respondWithJson(w, http.StatusOK, user)
}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	fmt.Println("Create New User : ", user)

	if err := daos.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, user)
}

func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := daos.Update(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := daos.Delete(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func CreateBankAccountOfUserEndPoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var userBankAcc model.BankAccout
	params := mux.Vars(r)

	user, err := daos.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&userBankAcc); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Println("userBankAcc :", userBankAcc)

	userBankAcc.ID = bson.NewObjectId()
	userBankAcc.UserID = user.ID

	user.BankAccount = append(user.BankAccount, userBankAcc)

	if err := daos.Update(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteBankAccountOfUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)
	user, err := daos.FindById(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	for _, element := range user.BankAccount {
		if element.ID == bson.ObjectIdHex(params["idAccount"]) {
			fmt.Println("Delete BankAccount.ID =", element.ID)
		} else {
			//user.BankAccount = user
			user.BankAccount = append(user.BankAccount, element)
			//ยังต้องแก้น้ะครับ
		}

	}

	if err := daos.Update(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
