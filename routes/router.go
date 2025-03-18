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
	docs.SwaggerInfo.BasePath = "/api/reef-metrics/v1"

	// Initialize the routes
	initializeRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Listen and serve on 0.0.0.0:8000
	router.Run(":8000")
}
