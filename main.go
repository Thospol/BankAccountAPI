package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AllUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "AllUserEndPoint !")
}

func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "FindUserEndpoint !")
}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CreateUserEndPoint !")
}

func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UpdateUserEndPoint !")
}

func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DeleteUserEndPoint !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", AllUserEndPoint).Methods("GET")
	r.HandleFunc("/users", FindUserEndpoint).Methods("POST")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", UpdateUserEndPoint).Methods("DELETE")
	r.HandleFunc("/users/{id}", DeleteUserEndPoint).Methods("GET")
	if err := http.ListenAndServe(":4525", r); err != nil {
		log.Fatal(err)
	}
}
