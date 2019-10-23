package database

import (
	"achievementsSystem/models"
	"errors"
)

// Database is used to get the "persistent" data
type Database struct {
	Users                map[string]models.User
	HistoricalStatistics map[string]models.HistoricalStatistics
}

var database *Database

// Init is used to initialize the database
func Init() error {
	historicalMap := make(map[string]models.HistoricalStatistics)
	usersMap := make(map[string]models.User)
	database = &Database{
		Users:                usersMap,
		HistoricalStatistics: historicalMap,
	}
	return nil
}

// LoadDummyUsers is used to load dummy users in dummy database
func LoadDummyUsers(users []models.User) error {
	if len(users) == 0 {
		return errors.New("Trying to init database with no users")
	}
	for _, user := range users {
		database.Users[user.ID] = user
	}
	return nil
}

// LoadDummyHistoricals is used to load dummy users in dummy database
func LoadDummyHistoricals(historicals []models.HistoricalStatistics) error {
	if len(historicals) == 0 {
		return errors.New("Trying to init database with driver containing no players historical statistics")
	}
	for _, historical := range historicals {
		database.HistoricalStatistics[historical.UserID] = historical
	}
	return nil
}
