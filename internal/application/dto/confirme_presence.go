package dto

type ConfirmePresenceDTO struct {
	Players []Players `json:"players"`
}

type Players struct {
	Name       string `json:"name"`
	Age        int    `json:"age,omitempty"`
	PlayerType string `json:"player_type"`
}
