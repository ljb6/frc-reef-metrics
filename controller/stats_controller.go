package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/frc-reef-metrics/models"
	"github.com/ljb6/frc-reef-metrics/usecase"
)

type StatsController struct {
	statsUsecase usecase.StatsUsecase
}

func NewStatsController(usecase usecase.StatsUsecase) StatsController {
	return StatsController{
		statsUsecase: usecase,
	}
}

func (r *StatsController) GetRows(ctx *gin.Context) {

	stats, err := r.statsUsecase.GetRows()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, stats)
}

func (r *StatsController) GetTeamData(ctx *gin.Context) {

	team := ctx.Param("team")
	fmt.Println(team)

	team_int, _ := strconv.Atoi(team)

	response := models.Response{
		Message: "Team number can't be null",
	}

	if team == " " {
		ctx.JSON(http.StatusInternalServerError, response)
	} else {
		stats, err := r.statsUsecase.GetTeamData(team_int)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, stats)
	}
}

func (r *StatsController) GetMatchData(ctx *gin.Context) {

	match := ctx.Param("match")

	match_int, _ := strconv.Atoi(match)

	response := models.Response{
		Message: "Match number can't be null",
	}

	if match == "" {
		ctx.JSON(http.StatusInternalServerError, response)
	} else {
		stats, err := r.statsUsecase.GetMatchData(match_int)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, stats)
	}
}