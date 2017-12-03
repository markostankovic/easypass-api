package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Email  string   `json:"email,omitempty"`
	Password  string   `json:"password,omitempty"`
}

func getUser(u *User) (db *sql.DB, err error) {
	statement := fmt.Sprintf("SELECT email FROM user WHERE email=?", u.Email)
	db.QueryRow(statement).Scan(u.Email)

	return
}

func GetUserById (userId string) (u *User, err error)  {
	u = &User{ ID: userId, Firstname: "Luka" }

	return
}

func NewUser(u *User) *User {
	return &User{ID: u.ID, Firstname: u.Firstname, }
}

func NewEmptyUser() *User {
	return &User{}
}
