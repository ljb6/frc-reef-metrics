package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljb6/frc-reef-metrics/controller"
	"github.com/ljb6/frc-reef-metrics/usecase"
)

func initializeRoutes(router *gin.Engine) {

	StatsUsecase := usecase.NewStatsUsecase()
	StatsController := controller.NewStatsController(StatsUsecase)

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