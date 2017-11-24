package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/markostankovic/easypass-api/models"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	Token string `json:"token"`
}

func GetUserById (w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user, err := models.GetUserById(params["id"])

	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(user)

	return
}

func GetAllUsers (w http.ResponseWriter, req *http.Request)  {
	users, err := models.GetAllUsers()

	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(users)
}

func NewUser (w http.ResponseWriter, req *http.Request)  {
	// params := mux.Vars(req)
	user := models.NewEmptyUser()
	json.NewDecoder(req.Body).Decode(&user)
	json.NewEncoder(w).Encode(user)
	fmt.Print(user)
}

func CreateToken (w http.ResponseWriter, req *http.Request)  {
	user := models.NewEmptyUser()
	_ = json.NewDecoder(req.Body).Decode(&user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"password": user.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}