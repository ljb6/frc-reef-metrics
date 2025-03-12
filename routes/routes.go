package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljb6/frc-reef-metrics/controller"
)

func initializeRoutes(router *gin.Engine) {

	StatsController := controller.NewStatsController()

	// Endpoint
	end := router.Group("/api/reef-metrics/v1")

	// Test
	end.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "pong",
		})
	})

	// Matches test
	end.GET("/matches", StatsController.GetMatches)
}