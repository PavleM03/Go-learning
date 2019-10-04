package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.Open("driver", "username:password@tcp(127.0.0.1:3306)/db")
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/test")
	// Error handling
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// db:= sql.Open -> db.Query("SQL QUERY")
	insert, err := db.Query("INSERT INTO tableroni(username,password) VALUES ('2','5') ")

	// Error handling
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
