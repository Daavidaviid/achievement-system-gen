// +build ignore

// This program generates multiple go files
// It can be invoked by running go generate

package main

import (
	generator "achievementsSystem/gen_files"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	defer fmt.Println("Generating the models, DONE!")
	rules, err := generator.GetRules()
	if err != nil {
		panic(err)
	}

	///////////
	// MODELS /
	///////////
	f, err := os.Create("models/models_gen.go")
	die(err)
	modelsTemplate := template.Must(template.New("models.tmpl").ParseFiles("gen_files/templates/models.tmpl"))
	die(err)
	modelsTemplate.Execute(f, rules)
	f.Close()

	///////////
	// DRIVER /
	///////////
	f, err = os.Create("driver/generateHistoricalStatistics_gen.go")
	die(err)
	generateHistoricalTemplate := template.Must(template.New("generateHistorical.tmpl").ParseFiles("gen_files/templates/generateHistorical.tmpl"))
	die(err)
	generateHistoricalTemplate.Execute(f, rules)
	f.Close()

	f, err = os.Create("driver/developGame_gen.go")
	die(err)
	developGameTemplate := template.Must(template.New("developGame.tmpl").ParseFiles("gen_files/templates/developGame.tmpl"))
	die(err)
	developGameTemplate.Execute(f, rules)
	f.Close()

	////////////////
	// ACHIEVEMENT /
	////////////////
	f, err = os.Create("achievementmanager/checkAchievements_gen.go")
	die(err)
	checkAchievementsTemplate := template.Must(template.New("checkAchievements.tmpl").ParseFiles("gen_files/templates/checkAchievements.tmpl"))
	die(err)
	checkAchievementsTemplate.Execute(f, rules)
	f.Close()

	f, err = os.Create("achievementmanager/getUserAchievements_gen.go")
	die(err)
	getUserAchievementsTemplate := template.Must(template.New("getUserAchievements.tmpl").ParseFiles("gen_files/templates/getUserAchievements.tmpl"))
	die(err)
	getUserAchievementsTemplate.Execute(f, rules)
	f.Close()

	f, err = os.Create("achievementmanager/processTeamAchievements_gen.go")
	die(err)
	processTeamAchievementsTemplate := template.Must(template.New("processTeamAchievements.tmpl").ParseFiles("gen_files/templates/processTeamAchievements.tmpl"))
	die(err)
	processTeamAchievementsTemplate.Execute(f, rules)
	f.Close()

	//////////
	// TESTS /
	//////////
	f, err = os.Create("achievementmanager/rules_gen_test.go")
	die(err)
	testRulesTemplate := template.Must(template.New("testRules.tmpl").ParseFiles("gen_files/templates/testRules.tmpl"))
	die(err)
	testRulesTemplate.Execute(f, rules)
	f.Close()

}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
