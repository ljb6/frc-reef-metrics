package models

type EventStats struct {
	Id           int    `json:"id"`
	Event_key    string `json:"event_key"`
	Week         string `json:"event_week"`
	Total_corals int    `json:"total_corals"`
}
