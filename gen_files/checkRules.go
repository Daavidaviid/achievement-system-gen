package generator

import (
	"fmt"
	"regexp"
	"strings"
)

// CheckRulesValidity is important because its role is to make sure the rules defined
// in the `rules.json` are legit
func CheckRulesValidity(rules *Rules) error {

	var gameStatisticsNames []string
	var historicalStatisticsNames []string

	varNameRegex, _ := regexp.Compile("(historical.|game.)\\w+")
	floatNumberRegex, _ := regexp.Compile("[0-9]+\\.[0-9]+")

	for _, gameStatistics := range rules.Statistics.Game {
		if gameStatistics.Type != "int" && gameStatistics.Type != "float" {
			return fmt.Errorf("Wrong type for %s, expected int or float, got %s", gameStatistics.Name, gameStatistics.Type)
		}
		gameStatisticsNames = append(gameStatisticsNames, gameStatistics.Name)
	}

	historicalStatisticsNames = gameStatisticsNames
	for _, historicalStatistics := range rules.Statistics.Historical {
		if historicalStatistics.Type != "int" && historicalStatistics.Type != "float" {
			return fmt.Errorf("Wrong type for %s, expected int or float, got %s", historicalStatistics.Name, historicalStatistics.Type)
		}
		historicalStatisticsNames = append(historicalStatisticsNames, historicalStatistics.Name)
	}

	for _, rule := range rules.Achievements.Rules {
		if floatNumberRegex.MatchString(rule.Rule) {
			return fmt.Errorf("No float numbers allowed in rule for the %s achievement", rule.Name)
		}
		for _, entity := range varNameRegex.FindAllString(rule.Rule, -1) {
			switch {
			case strings.HasPrefix(entity, "historical."):
				varName := strings.Split(entity, ".")[1]
				if !contains(historicalStatisticsNames, varName) {
					return fmt.Errorf("%s in rule expression <%s> for the %s achievement doesn't exist", entity, rule.Rule, rule.Key)
				}
				break
			case strings.HasPrefix(entity, "game."):
				varName := strings.Split(entity, ".")[1]
				if !contains(gameStatisticsNames, varName) {
					return fmt.Errorf("%s in rule expression <%s> for the %s achievement doesn't exist", entity, rule.Rule, rule.Key)
				}
				break

			}
		}

	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
