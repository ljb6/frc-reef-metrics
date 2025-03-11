package database

import (
	"database/sql"
	"log"

	"github.com/ljb6/frc-reef-metrics/models"
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
		week TEXT NOT NULL,
		auto_data TEXT NOT NULL,
		teleop_data TEXT NOT NULL
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration done")
}

func InsertData(data models.EventStats) {
	query := "INSERT INTO reef_stats (event_key, week, auto_data, teleop_data) VALUES (?, ?, ?, ?)"
	_, err := DB.Exec(query, data.Event_key, data.Week, data.Auto_data, data.Tele_data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data:", data, "added to DB")
}