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

	err = db.QueryRow(`SELECT user_id, email, firstname, lastname, active, company
		FROM user
		WHERE email=? and password=?`, u.Email, hash).Scan(
		&result.ID,
		&result.Email,
		&result.Firstname,
		&result.Lastname,
		&result.Active,
		&result.Company)

	if err != nil {
		fmt.Print(err.Error())
	}

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
	CREATE TABLE IF NOT EXISTS user (
		user_id int(11) NOT NULL AUTO_INCREMENT,
		username varchar(16) DEFAULT NULL,
		email varchar(255) NOT NULL,
		password varchar(255) NOT NULL,
		create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP,
		firstname varchar(45) DEFAULT NULL,
		lastname varchar(45) DEFAULT NULL,
		company varchar(45) DEFAULT NULL,
		active tinyint(4) DEFAULT NULL,
		PRIMARY KEY (user_id),
		UNIQUE KEY user_id_UNIQUE (user_id),
		UNIQUE KEY email_UNIQUE (email)
	) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8`

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
	const tableCreationQuery = `CREATE TABLE IF NOT EXISTS roles (
	  role_id int(11) NOT NULL,
	  name varchar(45) NOT NULL,
	  active tinyint(4) NOT NULL,
	  PRIMARY KEY (role_id),
	  UNIQUE KEY role_id_UNIQUE (role_id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8`
	_, err := db.Exec(tableCreationQuery)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func TestUserRoleTable()  {
	const tableCreationQuery = `CREATE TABLE IF NOT EXISTS user_role (
	  user_id int(11) NOT NULL,
	  role_id int(11) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	_, err := db.Exec(tableCreationQuery)
	if err != nil {
		fmt.Print(err.Error())
	}
}
