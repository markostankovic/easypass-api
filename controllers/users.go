package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/markostankovic/easypass-api/models"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"os"
)

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
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
	secret := os.Getenv("JWT_SECRET")
	tokenString, error := token.SignedString([]byte(secret))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					secret := os.Getenv("JWT_SECRET")
					return []byte(secret), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}
