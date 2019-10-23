package driver

import (
	"achievementsSystem/models"

	"github.com/bxcodec/faker"
)

// GenerateUsers is used to generate dummy users and return a slice containing all their IDs
func GenerateUsers(numberOfUsers int) ([]models.User, error) {
	var newUsers []models.User

	for index := 0; index < numberOfUsers; index++ {
		newUser := &models.User{}
		err := faker.FakeData(newUser)
		if err != nil {
			return []models.User{}, err
		}
		newUsers = append(newUsers, *newUser)
	}

	return newUsers, nil
}
