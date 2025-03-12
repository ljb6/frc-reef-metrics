package repository

import (
	"database/sql"
	"fmt"

	"github.com/ljb6/frc-reef-metrics/models"
)

type StatsRepository struct {
	conn *sql.DB
}

func NewStatsRepository(conn *sql.DB) StatsRepository {
	 return StatsRepository {
		conn: conn,
	 }
} 

func (sr *StatsRepository) GetProducts() ([]models.MatchStats, error) {
	
	query := "SELECT  FROM robot_match_stats"
	rows, err := sr.conn.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var stats []models.MatchStats
	var stat models.MatchStats

	for rows.Next() {
		err = rows.Scan(
			&stat.Name,
			&stat.Email,
			&stat.TeamNumber,
			&stat.MatchNumber,
			&stat.MatchLevel,
			&stat.StartZone,
	
			&stat.AutoLeft,
			&stat.AutoL1Corals,
			&stat.AutoL2Corals,
			&stat.AutoL3Corals,
			&stat.AutoL4Corals,
			&stat.AutoProcessor,
			&stat.AutoNet,
	
			&stat.TeleL1Corals,
			&stat.TeleL2Corals,
			&stat.TeleL3Corals,
			&stat.TeleL4Corals,
			&stat.TeleProcessor,
			&stat.TeleNet,
	
			&stat.EndPark,
			&stat.EndClimbAttempt,
			&stat.EndClimbLevel,
			&stat.EndClimbFailed,
	
			&stat.RemovedAlgae,
			&stat.RobotFailed,
			&stat.PlayedDefense,
			&stat.TrappedInAlgae,
			&stat.EndFouls,
			&stat.Comments,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		stats = append(stats, stat)
	}
	
	rows.Close()

	return stats, nil
}