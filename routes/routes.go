package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	// Endpoint
	end := router.Group("/api/reef-metrics/v1")

	// Test
	end.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "pong",
		})
	})
}