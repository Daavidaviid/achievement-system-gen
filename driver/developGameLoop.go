package driver

import "fmt"

// DevelopGameLoop is executing numberOfLoops times the method DevelopGame
func (driver *Driver) DevelopGameLoop(numberOfLoops int) error {
	defer fmt.Printf("Went through the develop game loop %d times.\n", numberOfLoops)
	var err error
	for index := 0; index < numberOfLoops; index++ {
		err = driver.DevelopGame()
		if err != nil {
			return err
		}
	}
	return nil
}
