package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Fatal("Failed to connect to SQLite:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS books (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create books table:", err)
	}

	log.Println("Database initialized and books table ensured.")
}
