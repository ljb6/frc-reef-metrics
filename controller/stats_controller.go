package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/frc-reef-metrics/models"
	"github.com/ljb6/frc-reef-metrics/usecase"
)

// StatsController gerencia as estatísticas da API
type StatsController struct {
	statsUsecase usecase.StatsUsecase
}

// NewStatsController cria uma nova instância do controlador de estatísticas
func NewStatsController(usecase usecase.StatsUsecase) StatsController {
	return StatsController{
		statsUsecase: usecase,
	}
}

// GetRows retorna todas as linhas de estatísticas
// @Summary Obtém todas as linhas de estatísticas
// @Description Retorna uma lista completa das estatísticas coletadas
// @Tags Stats
// @Produce json
// @Success 200 {array} models.MatchStats
// @Failure 500 {object} models.Response
// @Router /stats [get]
func (r *StatsController) GetRows(ctx *gin.Context) {
	stats, err := r.statsUsecase.GetRows()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, stats)
}

// GetTeamData retorna os dados estatísticos de um time específico
// @Summary Obtém os dados de um time
// @Description Retorna as estatísticas de um time específico pelo número
// @Tags Stats
// @Produce json
// @Param team path int true "Número do time"
// @Success 200 {object} models.MatchStats
// @Failure 400 {object} models.Response "Erro: Número do time inválido"
// @Failure 500 {object} models.Response "Erro interno"
// @Router /stats/team/{team} [get]
func (r *StatsController) GetTeamData(ctx *gin.Context) {
	team := ctx.Param("team")

	notNumResponse := models.Response{
		Message: "Team number needs to be an integer.",
	}
	noDataResponse := models.Response{
		Message: "There is no data for this team.",
	}

	_, err := strconv.Atoi(team)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notNumResponse)
		return
	}

	stats, err := r.statsUsecase.GetTeamData(team)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if stats == nil {
		ctx.JSON(http.StatusNotFound, noDataResponse)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

// GetMatchData retorna os dados estatísticos de uma partida específica
// @Summary Obtém os dados de uma partida
// @Description Retorna as estatísticas de uma partida pelo número da partida
// @Tags Stats
// @Produce json
// @Param match path int true "Número da partida"
// @Success 200 {object} models.MatchStats
// @Failure 400 {object} models.Response "Erro: Número da partida inválido"
// @Failure 500 {object} models.Response "Erro interno"
// @Router /stats/match/{match} [get]
func (r *StatsController) GetMatchData(ctx *gin.Context) {
	match := ctx.Param("match")

	notNumResponse := models.Response{
		Message: "Match number needs to be an integer.",
	}
	noDataResponse := models.Response{
		Message: "There is no data for this match.",
	}

	_, err := strconv.Atoi(match)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notNumResponse)
		return
	}

	stats, err := r.statsUsecase.GetMatchData(match)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if stats == nil {
		ctx.JSON(http.StatusNotFound, noDataResponse)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

// GetTeamMatch retorna os dados estatísticos de um time em uma partida específica
// @Summary Obtém os dados de um time em uma partida específica
// @Description Retorna as estatísticas de um time dentro de uma determinada partida
// @Tags Stats
// @Produce json
// @Param match path int true "Número da partida"
// @Param team path int true "Número do time"
// @Success 200 {object} models.MatchStats
// @Failure 400 {object} models.Response "Erro: Parâmetros inválidos"
// @Failure 500 {object} models.Response "Erro interno"
// @Router /stats/match/{match}/team/{team} [get]
func (r *StatsController) GetTeamMatch(ctx *gin.Context) {
	match := ctx.Param("match")
	team := ctx.Param("team")

	noDataResponse := models.Response{
		Message: "There is no data for this team and match.",
	}

	stats, err := r.statsUsecase.GetTeamMatch(match, team)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if stats == nil {
		ctx.JSON(http.StatusNotFound, noDataResponse)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}
