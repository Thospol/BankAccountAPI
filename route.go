package main

import (
	"log"
	"net/http"
)

func InitRoute() {
	r.HandleFunc("/api/users", AllUserEndPoint).Methods("GET")
	r.HandleFunc("/api/users/{id}", FindUserEndpoint).Methods("GET")
	r.HandleFunc("/api/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/api/users", UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/api/users", DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/api/users/{id}/bankAccount", CreateBankAccountOfUserEndPoint).Methods("POST")

	if err := http.ListenAndServe(":4525", r); err != nil {
		log.Fatal(err)
	}
}
