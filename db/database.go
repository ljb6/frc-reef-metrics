package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./stats.db")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database oppened")
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
		log.Println("Database closed")
	}
}

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS reef_stats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_key TEXT NOT NULL,
		week, TEXT NOT NULL,
		total_corals INTEGER NOT NULL,
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration done")
}