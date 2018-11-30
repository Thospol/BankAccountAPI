package main

import (
	"log"
	"net/http"
)

func InitRoute() {
	r.HandleFunc("/users", AllUserEndPoint).Methods("GET")
	r.HandleFunc("/users", FindUserEndpoint).Methods("POST")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", UpdateUserEndPoint).Methods("DELETE")
	r.HandleFunc("/users/{id}", DeleteUserEndPoint).Methods("GET")

	if err := http.ListenAndServe(":4525", r); err != nil {
		log.Fatal(err)
	}
}
