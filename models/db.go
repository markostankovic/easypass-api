package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	// defer db.Close()

	if err != nil {
		fmt.Print(err.Error())
	}

	shouldTestDB := os.Getenv("TEST_DB")

	if len(shouldTestDB) > 0 {
		TestDB()
	}
}
