package models

// Team represents a team during a game
type Team struct {
	ID      string             `json:"id"`
	Score   int                `json:"score"`
	Players []GameStatistics `json:"players"`
}
