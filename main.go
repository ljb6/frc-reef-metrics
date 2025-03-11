package main

import (
	database "github.com/ljb6/frc-reef-metrics/db"
	"github.com/ljb6/frc-reef-metrics/models"
	router "github.com/ljb6/frc-reef-metrics/routes"
)

func main() {
	database.OpenDatabase()
	defer database.CloseDatabase()

	database.Migrate()

	test_data := models.EventStats{
		Event_key: "brba2024",
		Week: "Week 3",
		Auto_data: `"botRow": { "nodeA": 3, "nodeB": 5, "nodeC": 2 },
					"midRow": { "nodeA": 4, "nodeB": 1, "nodeC": 3 },
					"topRow": { "nodeA": 2, "nodeB": 3, "nodeC": 1 },
					"trough": 0`,
		Tele_data: `"botRow": { "nodeA": 3, "nodeB": 5, "nodeC": 2 },
					"midRow": { "nodeA": 4, "nodeB": 1, "nodeC": 3 },
					"topRow": { "nodeA": 2, "nodeB": 3, "nodeC": 1 },
					"trough": 0`,
	}

	database.InsertData(test_data)

	router.InitializeServer()
}