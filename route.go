package main

import (
	"log"
	"net/http"
)

func SetupRoute() {
	r.HandleFunc("/api/v1/users", AllUserEndPoint).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", FindUserEndpoint).Methods("GET")
	r.HandleFunc("/api/v1/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/api/v1/users", UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/api/v1/users", DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/api/v1/users/{id}/bankAccount", CreateBankAccountOfUserEndPoint).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}/bankAccount", FindUserEndpoint).Methods("GET")
	r.HandleFunc("/api/v1/bankAccount/{id}/{idAccount}", DeleteBankAccountOfUsersEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":4525", r); err != nil {
		log.Fatal(err)
	}
}
