package main

import (
	"achievementsSystem/database"
	"achievementsSystem/driver"
	"fmt"
	"log"
	"math/rand"
	"time"
)

//go:generate go run gen/gen.go

const (
	numberOfUsersInTheSystem = 15
	teamSize                 = 4
	numberOfLoop             = 5
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("## [TECHNICAL TEST] Achievements system for a fictional online game ##")

	currentDriver, err := driver.New(numberOfUsersInTheSystem)
	die(err, "init driver")

	err = database.Init()
	die(err, "init database")

	err = database.LoadDummyUsers(currentDriver.Users)
	die(err, "load dummy users in database")

	err = database.LoadDummyHistoricals(currentDriver.HistoricalStatistics)
	die(err, "load dummy historicals in database")

	log.Println("PROUT1")
	err = currentDriver.CreateGame(teamSize)
	die(err, "create game")

	log.Println("PROUT2")

	err = currentDriver.DevelopGameLoop(numberOfLoop)
	die(err, "develop game")

	err = currentDriver.EndGame()
	die(err, "end game")
}

func die(err error, name string) {
	if err != nil {
		log.Fatal(name, ":", err)
	}
}
