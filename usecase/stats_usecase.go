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

func (su *StatsUsecase) GetMatches() ([]models.MatchStats, error) {
	return nil, nil
}