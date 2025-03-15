package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/frc-reef-metrics/controller"
	"github.com/ljb6/frc-reef-metrics/db"
	"github.com/ljb6/frc-reef-metrics/repository"
	"github.com/ljb6/frc-reef-metrics/usecase"
)

func initializeRoutes(router *gin.Engine) {

	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	StatsRepository := repository.NewStatsRepository(dbConnection)
	StatsUseCase := usecase.NewStatsUsecase(StatsRepository)
	StatsController := controller.NewStatsController(StatsUseCase)

	// Endpoint
	end := router.Group("/api/reef-metrics/v1")

	// Test
	end.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "pong",
		})
	})

	end.GET("/matches", StatsController.GetRows)
	end.GET("/team/:team", StatsController.GetTeamData)
}