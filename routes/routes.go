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

	// Endpoints
	end := router.Group("/api/reef-metrics/v1")

	end.GET("/all-matches", StatsController.GetRows)
	end.GET("/matches/:team", StatsController.GetTeamData)
	end.GET("/match/:match", StatsController.GetMatchData)
	end.GET("/match/:match/:team", StatsController.GetTeamMatch)
	
}