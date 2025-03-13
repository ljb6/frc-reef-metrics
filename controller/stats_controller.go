package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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