package usecase

import (
	"log"

	"github.com/ljb6/frc-reef-metrics/models"
	"github.com/ljb6/frc-reef-metrics/repository"
)

type StatsUsecase struct {
	repository repository.StatsRepository
}

func NewStatsUsecase(repo repository.StatsRepository) StatsUsecase {
	return StatsUsecase{
		repository: repo,
	}
}

func (su *StatsUsecase) GetRows() ([]models.MatchStats, error) {
	log.Println("Searching for all matches...")
	return su.repository.GetRows()
}

func (su *StatsUsecase) GetTeamData(team string) ([]models.MatchStats, error) {
	return su.repository.GetTeamData(team)
}

func (su *StatsUsecase) GetMatchData(match string) ([]models.MatchStats, error) {
	return su.repository.GetMatchData(match)
}