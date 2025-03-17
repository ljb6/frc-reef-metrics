package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/ljb6/frc-reef-metrics/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

func InitializeServer() {
	// Iinitialize router
	router := gin.Default()

	// Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Initialize the routes
	initializeRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
