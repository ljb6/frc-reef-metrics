package main

import database "github.com/ljb6/frc-reef-metrics/db"

func main() {
	database.OpenDatabase()
	defer database.CloseDatabase()

	database.Migrate()
}