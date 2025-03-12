package usecase

import "github.com/ljb6/frc-reef-metrics/models"

type statsUsecase struct {
	//
}

func NewStatsUsecase() statsUsecase {
	return statsUsecase{}
}

func (su *statsUsecase) GetMatches() ([]models.MatchStats, error) {
	return nil, nil
}