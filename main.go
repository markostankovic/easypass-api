package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/markostankovic/easypass-api/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user", controllers.NewUser).Methods("POST")
	router.HandleFunc("/authenticate", controllers.CreateToken).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}