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
	return StatsRepository{
		conn: conn,
	}
}

func scanMatchStats(rows *sql.Rows) ([]models.MatchStats, error) {

	var stats []models.MatchStats

	for rows.Next() {
		var stat models.MatchStats
		err := rows.Scan(
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
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

func (sr *StatsRepository) GetRows() ([]models.MatchStats, error) {

	query := "SELECT name, email, team_number, match_number, match_level, start_zone, " +
		"auto_left, auto_l1_corals, auto_l2_corals, auto_l3_corals, auto_l4_corals, auto_processor, auto_net, " +
		"tele_l1_corals, tele_l2_corals, tele_l3_corals, tele_l4_corals, tele_processor, tele_net, " +
		"end_park, end_climb_attempt, end_climb_level, end_climb_failed, " +
		"removed_algae, robot_failed, played_defense, trapped_in_algae, end_fouls " +
		"FROM robot_match_stats"

	rows, err := sr.conn.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	return scanMatchStats(rows)
}

func (sr *StatsRepository) GetTeamData(team int) ([]models.MatchStats, error) {

	query := `
	SELECT name, email, team_number, match_number, match_level, start_zone,
		auto_left, auto_l1_corals, auto_l2_corals, auto_l3_corals, auto_l4_corals, auto_processor, auto_net,
		tele_l1_corals, tele_l2_corals, tele_l3_corals, tele_l4_corals, tele_processor, tele_net,
		end_park, end_climb_attempt, end_climb_level, end_climb_failed,
		removed_algae, robot_failed, played_defense, trapped_in_algae, end_fouls
	FROM robot_match_stats 
	WHERE team_number = $1`

	rows, err := sr.conn.Query(query, team)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	return scanMatchStats(rows)
}

func (sr *StatsRepository) GetMatchData(match int) ([]models.MatchStats, error) {

	query := `
	SELECT name, email, team_number, match_number, match_level, start_zone,
		auto_left, auto_l1_corals, auto_l2_corals, auto_l3_corals, auto_l4_corals, auto_processor, auto_net,
		tele_l1_corals, tele_l2_corals, tele_l3_corals, tele_l4_corals, tele_processor, tele_net,
		end_park, end_climb_attempt, end_climb_level, end_climb_failed,
		removed_algae, robot_failed, played_defense, trapped_in_algae, end_fouls
	FROM robot_match_stats 
	WHERE match_number = $1`

	rows, err := sr.conn.Query(query, match)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	return scanMatchStats(rows)
}