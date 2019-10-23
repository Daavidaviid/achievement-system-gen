package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func getRulesFileContent() (*os.File, error) {
	jsonFile, err := os.Open("rules.json")
	if err != nil {
		return nil, err
	}
	return jsonFile, err
}

// GetRules is used to get the rules from the rules.json file
func GetRules() (*Rules, error) {
	file, err := getRulesFileContent()
	defer file.Close()
	if err != nil {
		return nil, err
	}
	rulesBytes, _ := ioutil.ReadAll(file)

	rules := &Rules{}
	json.Unmarshal(rulesBytes, rules)

	err = CheckRulesValidity(rules)

	varNameRegex, _ := regexp.Compile("(historical.|game.)\\w+")
	for index, rule := range rules.Achievements.Rules {
		formattedRule := varNameRegex.ReplaceAllStringFunc(rule.Rule, func(m string) string {
			return fmt.Sprintf("float64(%s)", m)
		})
		rules.Achievements.Rules[index].FormattedRule = formattedRule
	}

	return rules, err
}
