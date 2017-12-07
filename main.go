package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/markostankovic/easypass-api/controllers"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/markostankovic/easypass-api/models"
	"os"
)

func main() {
	fmt.Println("Starting the application...")
	dataAcceess := os.Getenv("DATA_ACCESS") + "@/easypass"

	models.InitDB(dataAcceess)
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.ValidateMiddleware(controllers.GetAllUsers)).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.ValidateMiddleware(controllers.GetUserById)).Methods("GET")
	router.HandleFunc("/user", controllers.ValidateMiddleware(controllers.NewUser)).Methods("POST")
	router.HandleFunc("/login", controllers.CreateToken).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}