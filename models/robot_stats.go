package models

type RobotMatchStats struct {
	// Pre match
	Name            string `json:"name"`
	Email           string `json:"email"`
	TeamNumber      int    `json:"team_number"`
	MatchNumber     int    `json:"match_number"`
	MatchLevel      string `json:"match_level"`
	StartZone       string `json:"start_zone"`

	// Autonomous
	AutoLeft        bool   `json:"auto_left"`
	AutoL1Corals    int    `json:"auto_l1_corals"`
	AutoL2Corals    int    `json:"auto_l2_corals"`
	AutoL3Corals    int    `json:"auto_l3_corals"`
	AutoL4Corals    int    `json:"auto_l4_corals"`
	AutoProcessor   int   `json:"auto_processor"`
	AutoNet         int    `json:"auto_net"`

	// Teleop
	TeleL1Corals    int    `json:"tele_l1_corals"`
	TeleL2Corals    int    `json:"tele_l2_corals"`
	TeleL3Corals    int    `json:"tele_l3_corals"`
	TeleL4Corals    int    `json:"tele_l4_corals"`
	TeleProcessor   int   `json:"tele_processor"`
	TeleNet         int    `json:"tele_net"`

	// End
	EndPark         bool   `json:"end_park"`
	EndClimbAttempt bool   `json:"end_climb_attempt"`
	EndClimbLevel   string    `json:"end_climb_level"`
	EndClimbFailed  bool   `json:"end_climb_failed"`

	// Both
	RemovedAlgae    bool    `json:"removed_algae"`
	RobotFailed     bool   `json:"robot_failed"`
	PlayedDefense   bool   `json:"played_defense"`
	TrappedInAlgae  bool   `json:"trapped_in_algae"`
	EndFouls        int    `json:"end_fouls"`
	Comments        string `json:"comments"`
}
