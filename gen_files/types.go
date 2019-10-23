package generator

type Rules struct {
	Statistics   Statistics   `json:"statistics"`
	Achievements Achievements `json:"achievements"`
}

type Statistics struct {
	Game       []Statistic `json:"game"`
	Historical []Statistic `json:"historical"`
}

type Statistic struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Achievements struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Name          string `json:"name"`
	Key           string `json:"key"`
	Description   string `json:"description"`
	ApplyTo       string `json:"applyTo"`
	Rule          string `json:"rule"`
	FormattedRule string `json:"formattedRule"`
	Tests         *Tests `json:"tests"`
}

type Tests struct {
	Success *Scenario `json:"success"`
	Failure *Scenario `json:"failure"`
}

type Scenario struct {
	Game       *map[string]int `json:"game"`
	Historical *map[string]int `json:"historical"`
}
