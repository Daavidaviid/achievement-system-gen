package models

// Game contains duration, numberOfPlayer, and teams data
type Game struct {
	ID       string `json:"id"`
	TeamSize int    `json:"teamSize"`
	BlueTeam Team   `json:"blueTeam"`
	RedTeam  Team   `json:"redTeam"`
}
