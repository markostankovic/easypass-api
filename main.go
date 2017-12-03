package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/markostankovic/easypass-api/controllers"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	//db, err = sql.Open("mysql", "marenevreme:1Carolija@/easypass")
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err.Error())
	//}
}

func (a *App) Run(addr string) {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.ValidateMiddleware(controllers.GetAllUsers)).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.ValidateMiddleware(controllers.GetUserById)).Methods("GET")
	router.HandleFunc("/user", controllers.ValidateMiddleware(controllers.NewUser)).Methods("POST")
	router.HandleFunc("/login", controllers.CreateToken).Methods("POST")
	log.Fatal(http.ListenAndServe(addr, router))
}

func main() {
	fmt.Println("Starting the application...")

	a := App{}
	a.Initialize("marenevreme", "1Carolija", "easypass")

	a.Run(":12345")
}