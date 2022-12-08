package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miily8310s/crud-rest-api/user"
)

func initializeRouter()  {
	r := mux.NewRouter()
	// r.HandleFunc("/users", GetUsers).Methods("GET")
	// r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", user.CreateUser).Methods("POST")
	// r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	// r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	user.InitializeMigration()
	initializeRouter()
}