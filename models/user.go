package models

// User contains informations
type User struct {
	ID                string `json:"id" faker:"uuid_hyphenated"`
	Name              string `json:"name" faker:"name"`
	IsAlreadySelected bool
}
