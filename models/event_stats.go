package models

type EventStats struct {
	// Event info
	Id        int    `json:"id"`
	Event_key string `json:"event_key"`
	Week      string `json:"event_week"`

	// Periods data
	Auto_data string `json:"auto_data"`
	Tele_data string `json:"teleop_data"`
}
