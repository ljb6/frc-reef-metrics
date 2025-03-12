package main

import router "github.com/ljb6/frc-reef-metrics/routes"

func main() {

	// dbConnection, err := db.ConnectDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	router.InitializeServer()
}