package achievementmanager

import (
	"achievementsSystem/models"
)

// ProcessAchievements is called after a game is done to compute achievements for each player
func ProcessAchievements(endedGame *models.Game, historicals []models.HistoricalStatistics) ([]models.HistoricalStatistics, error) {
	// areTeamExaequo := endedGame.BlueTeam.Score == endedGame.RedTeam.Score
	didBlueTeamWon := endedGame.BlueTeam.Score > endedGame.RedTeam.Score

	processTeamAchievements(endedGame.BlueTeam.Players, &historicals, didBlueTeamWon)
	processTeamAchievements(endedGame.RedTeam.Players, &historicals, !didBlueTeamWon)

	return historicals, nil
}
