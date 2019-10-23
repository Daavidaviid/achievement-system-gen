package driver

import (
	"achievementsSystem/models"
	"fmt"
)

// Driver is used to test the app
type Driver struct {
	Users                []models.User
	HistoricalStatistics []models.HistoricalStatistics
	Game                 *models.Game
}

// New is used to create a new Driver
func New(numberOfUsers int) (*Driver, error) {
	var driver Driver
	users, err := GenerateUsers(numberOfUsers)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Generated %d dummy users.\n", numberOfUsers)

	historicalStatistics, err := GenerateHistoricalStatistics(users)
	if err != nil {
		return nil, err
	}
	fmt.Println("Generated dummy historical statistics for each user.")

	driver.Users = users
	driver.HistoricalStatistics = historicalStatistics

	return &driver, nil
}

// ClearData is used to empty the driver data once database has been filled
func (driver *Driver) ClearData() {
	driver.Users = nil
	driver.HistoricalStatistics = nil
}
