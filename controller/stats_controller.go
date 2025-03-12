package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/frc-reef-metrics/models"
)

type robotStatsController struct {
	// 
}

func NewRobotStatsController() robotStatsController {
	return robotStatsController{}
}

func (r *robotStatsController) GetMatches(ctx *gin.Context) {
	stats := []models.RobotMatchStats{
		{
			Name:        "John Doe",
			Email:       "john.doe@example.com",
			TeamNumber:  1234,
			MatchNumber: 5,
			MatchLevel:  "Qualification",
			StartZone:   "Blue",
			AutoLeft:    true,
			AutoL1Corals: 2,
			AutoL2Corals: 1,
			AutoL3Corals: 0,
			AutoL4Corals: 0,
			AutoProcessor: 1,
			AutoNet:      3,
			TeleL1Corals: 4,
			TeleL2Corals: 2,
			TeleL3Corals: 1,
			TeleL4Corals: 0,
			TeleProcessor: 3,
			TeleNet:      5,
			EndPark:      true,
			EndClimbAttempt: true,
			EndClimbLevel:   "High",
			EndClimbFailed:  false,
			RemovedAlgae:    true,
			RobotFailed:     false,
			PlayedDefense:   true,
			TrappedInAlgae:  false,
			EndFouls:        2,
			Comments:        "Good match, strong defense",
		},
	}

	ctx.JSON(http.StatusOK, stats)
}