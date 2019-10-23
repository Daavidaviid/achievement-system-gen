package database

import (
	"achievementsSystem/models"
	"fmt"
)

// GetHistoricalStatisticsByUserIDs returns a slice of models.HistoricalStatistics for the given IDs
func GetHistoricalStatisticsByUserIDs(usersIDs []string) ([]models.HistoricalStatistics, error) {
	var historicalStatistics []models.HistoricalStatistics
	for _, userID := range usersIDs {
		playerHistoricalStatistics, ok := database.HistoricalStatistics[userID]
		if !ok {
			return nil, fmt.Errorf("User's historical statistics not found for this specific userID : %s", userID)
		}
		historicalStatistics = append(historicalStatistics, playerHistoricalStatistics)
	}

	return historicalStatistics, nil
}
