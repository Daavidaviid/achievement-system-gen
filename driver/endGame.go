package driver

import (
	"achievementsSystem/achievementmanager"
	"achievementsSystem/database"
	"errors"
	"fmt"
)

// EndGame triggers the end of the game
func (driver *Driver) EndGame() error {
	// TODO: Send game to Achievements manager
	// game.
	if driver.Game == nil {
		return errors.New("Can't end game because no game exists")
	}

	endedGame := driver.Game

	// Getting players UserIDs
	var involvedUsersIDs []string
	for _, gameStatistics := range endedGame.BlueTeam.Players {
		involvedUsersIDs = append(involvedUsersIDs, gameStatistics.UserID)
	}
	for _, gameStatistics := range endedGame.RedTeam.Players {
		involvedUsersIDs = append(involvedUsersIDs, gameStatistics.UserID)
	}

	// Getting game players historical statistics
	historicals, err := database.GetHistoricalStatisticsByUserIDs(involvedUsersIDs)
	if err != nil {
		return err
	}

	historicals, err = achievementmanager.ProcessAchievements(endedGame, historicals)
	if err != nil {
		return err
	}

	fmt.Println("Displaying game achievements for each game players : ")
	for _, historical := range historicals {
		fmt.Printf("User (%s) has : \n\tAchievements : %s\n\tJust unlocked achievements : %s\n",
			historical.UserID,
			historical.Achievements,
			historical.NewAchievements,
		)
	}

	// TODO: We should write the historical statistics back into the database

	return nil
}
