package main

import (
	"log"
	"net/http"
)

func InitRoute() {
	r.HandleFunc("/users", AllUserEndPoint).Methods("GET")
	r.HandleFunc("/users/{id}", FindUserEndpoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users", UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", DeleteUserEndPoint).Methods("DELETE")

	if err := http.ListenAndServe(":4525", r); err != nil {
		log.Fatal(err)
	}
}
