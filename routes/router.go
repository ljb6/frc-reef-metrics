package router

import "github.com/gin-gonic/gin"

func InitializeServer() {
	// Iinitialize router
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}