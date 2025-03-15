package usecase

import (
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
	return su.repository.GetRows() 
}

func (su *StatsUsecase) GetTeamData(team int) ([]models.MatchStats, error) {
	return su.repository.GetTeamData(team) 
}		