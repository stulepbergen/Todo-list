package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DataBase() *sql.DB {
	db, _ := sql.Open("sqlite3", "./database.db")

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS todolist (id INTEGER PRIMARY KEY, item TEXT, done BOOLEAN DEFAULT FALSE)")

	if err != nil {
		fmt.Println(err)
	}

	return db
}
