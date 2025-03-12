package usecase

import "github.com/ljb6/frc-reef-metrics/models"

type StatsUsecase struct {
	//
}

func NewStatsUsecase() StatsUsecase {
	return StatsUsecase{}
}

func (su *StatsUsecase) GetMatches() ([]models.MatchStats, error) {
	return nil, nil
}