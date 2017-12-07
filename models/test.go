package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func TestDB() {
	fmt.Println("Testing application...")

	TestUserTable()
	TestRolesTable()
	TestUserRoleTable()
}