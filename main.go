package main

import (
	database "github.com/ljb6/frc-reef-metrics/db"
	"github.com/ljb6/frc-reef-metrics/models"
)

func main() {
	database.OpenDatabase()
	defer database.CloseDatabase()

	database.Migrate()

	test_data := models.EventStats{
		Event_key: "brba2024",
		Week: "Week 3",
		Total_corals: 432,
	}

	database.InsertData(test_data)
}