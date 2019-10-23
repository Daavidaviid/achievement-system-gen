package driver

import (
	"achievementsSystem/models"
	"errors"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

// CreateGame is used to create a dummy game
func (driver *Driver) CreateGame(teamSize int) error {
	defer fmt.Printf("Created a game with two %d players teams.\n", teamSize)
	newGameUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	newGame := &models.Game{
		ID:       newGameUUID.String(),
		TeamSize: teamSize,
	}

	blueTeam, err := driver.createTeam(teamSize)
	if err != nil {
		return err
	}
	redTeam, err := driver.createTeam(teamSize)
	if err != nil {
		return err
	}

	newGame.BlueTeam = *blueTeam
	newGame.RedTeam = *redTeam

	driver.Game = newGame

	return nil
}

func (driver *Driver) createTeam(teamSize int) (*models.Team, error) {
	newTeamUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	newTeamPlayers, err := driver.pickTeamPlayers(teamSize)
	if err != nil {
		return nil, err
	}
	newTeam := &models.Team{
		ID:      newTeamUUID.String(),
		Players: newTeamPlayers,
	}
	return newTeam, nil
}

func (driver *Driver) pickTeamPlayers(teamSize int) ([]models.GameStatistics, error) {
	var teamPlayers []models.GameStatistics
	if len(driver.Users) == 0 {
		return nil, errors.New("No users found")
	}
	if teamSize < 3 || teamSize > 5 {
		return nil, errors.New("Specify a number of players in-between 3 and 5")
	}
	if len(driver.Users) < teamSize {
		return nil, errors.New("Not enough players")
	}

	var usersIDsPool []string

	// Creating a dummy pool of users willing to play
	for _, user := range driver.Users {
		usersIDsPool = append(usersIDsPool, user.ID)
	}

	numberOfMissingPlayers := teamSize
	for numberOfMissingPlayers != 0 {
		selectedUserIndex := rand.Intn(len(usersIDsPool))
		userID := usersIDsPool[selectedUserIndex]

		// This condition is to make sure the function will stop at some point
		if len(usersIDsPool) == 0 {
			return nil, errors.New("Not enough users")
		}

		newGameStatistics := models.GameStatistics{
			UserID: userID,
		}

		teamPlayers = append(teamPlayers, newGameStatistics)
		usersIDsPool = append(usersIDsPool[:selectedUserIndex], usersIDsPool[selectedUserIndex+1:]...)

		numberOfMissingPlayers--
	}

	// log.Printf("teamPlayers: %#+v\n", teamPlayers)

	return teamPlayers, nil
}
