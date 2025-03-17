package controller

import (
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

	notNumResponse := models.Response{
		Message: "Team number needs to be a integer.",
	}
	noDataResponse := models.Response{
		Message: "There is no data of this team.",
	}

	_, err := strconv.Atoi(team)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notNumResponse)
		return
	}

	stats, err := r.statsUsecase.GetTeamData(team)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if stats == nil {
		ctx.JSON(http.StatusInternalServerError, noDataResponse)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

func (r *StatsController) GetMatchData(ctx *gin.Context) {

	match := ctx.Param("match")

	notNumResponse := models.Response{
		Message: "Match number needs to be a integer.",
	}
	noDataResponse := models.Response{
		Message: "There is no data of this team.",
	}

	_, err := strconv.Atoi(match)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notNumResponse)
		return
	}

	stats, err := r.statsUsecase.GetMatchData(match)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if stats == nil {
		ctx.JSON(http.StatusInternalServerError, noDataResponse)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

func (r *StatsController) GetTeamMatch(ctx *gin.Context) {
	match := ctx.Param("match")
	team := ctx.Param("team")

	noDataResponse := models.Response{
		Message: "There is no data of this team and match.",
	}

	stats, err := r.statsUsecase.GetTeamMatch(match, team)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if stats == nil {
		ctx.JSON(http.StatusInternalServerError, noDataResponse)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}