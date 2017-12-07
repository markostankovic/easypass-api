package models

import (
	"encoding/hex"
	"crypto/sha256"
	"fmt"
)

type User struct {
	ID        	string   `json:"user_id,omitempty"`
	CreateTime 	string   `json:"create_time,omitempty"`
	Company 	string   `json:"company,omitempty"`
	Active 		string   `json:"active,omitempty"`
	Firstname 	string   `json:"firstname,omitempty"`
	Lastname  	string   `json:"lastname,omitempty"`
	Email		string   `json:"email,omitempty"`
	Password  	string   `json:"password,omitempty"`
}

func GetUser(u *User) (nu *User, err error) {
	result := &User{}
	h := sha256.New()
	h.Write([]byte(u.Password))
	hash := hex.EncodeToString(h.Sum(nil))

	db.QueryRow("SELECT user_id, email, firstname, lastname, active, company FROM user WHERE email=? and password=?", u.Email, hash).Scan(&result.ID, &result.Email, &result.Firstname, &result.Lastname, &result.Active, &result.Company)

	return result, err
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

func TestUserTable()  {
	const tableCreationQuery = `
	CREATE TABLE IF NOT EXISTS easypass.user (
	  user_id INT NOT NULL,
	  username VARCHAR(16) NULL,
	  email VARCHAR(255) NOT NULL,
	  password VARCHAR(255) NOT NULL,
	  create_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	  firstname VARCHAR(45) NULL,
	  lastname VARCHAR(45) NULL,
	  company VARCHAR(45) NULL,
	  active  TINYINT NULL,
	  UNIQUE INDEX user_id_UNIQUE (user_id ASC),
	  PRIMARY KEY (user_id),
	  UNIQUE INDEX email_UNIQUE (email ASC))`

	const insertUser = "INSERT INTO user(username, email, password, firstname, lastname, active) VALUES(?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(tableCreationQuery)
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = db.Exec(insertUser,
		"marenevreme",
		"markostankovic87@gmail.com",
		"23abb2abadd90bfb7d30d9cea46cf902c8d0e2a3df12b339e25902b9164a422a",
		"Marko",
		"Stankovic",
		"1")
	if err != nil {
		fmt.Print(err.Error())
	}
}

func TestRolesTable()  {
	const tableCreationQuery = `CREATE TABLE IF NOT EXISTS easypass.roles (
  role_id INT NOT NULL,
  name VARCHAR(45) NOT NULL,
  active TINYINT NOT NULL,
  PRIMARY KEY (role_id),
  UNIQUE INDEX role_id_UNIQUE (role_id ASC))`
	_, err := db.Exec(tableCreationQuery)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func TestUserRoleTable()  {
	const tableCreationQuery = `CREATE TABLE IF NOT EXISTS easypass.user_role (
	  user_id INT NOT NULL,
	  role_id INT NOT NULL,
	  INDEX role_id_idx (role_id ASC),
	  CONSTRAINT user_id
		FOREIGN KEY (user_id)
		REFERENCES easypass.user (user_id)
		ON DELETE NO ACTION
		ON UPDATE NO ACTION,
	  CONSTRAINT role_id
		FOREIGN KEY (role_id)
		REFERENCES easypass.roles (role_id)
		ON DELETE NO ACTION
		ON UPDATE NO ACTION)`

	_, err := db.Exec(tableCreationQuery)
	if err != nil {
		fmt.Print(err.Error())
	}
}
