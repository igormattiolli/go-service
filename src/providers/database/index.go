package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabaseServer() *sql.DB {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	return db
}

func CloseDatabaseServer(db *sql.DB) {
	defer db.Close()
}
