# Thought process during the development

At first my GameStatistics was like that :

```
type GameStatistics struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	GameID string `json:"gameId"`
	TeamID string `json:"teamId"`

	NumberOfAttacks       int `json:"numberOfAttacks"`
	NumberOfHits          int `json:"numberOfHits"`
	TotalAmountDamageDone int `json:"totalAmountDamageDone"`
	NumberOfKills         int `json:"numberOfKills"`
	NumberOfFirstHitKills int `json:"numberOfFirstHitKills"`
	NumberOfAssists       int `json:"numberOfAssists"`
	TotalNumberSpellsCast int `json:"totalNumberSpellsCast"`
	TotalSpellDamageDone  int `json:"totalSpellDamageDone"`
	TotalTimePlayed       int `json:"totalTimePlayed"`
}
```

It was nice but definitely not flexible enough to add new stats, achivements, etc.
So, I got the idea to use codegen, since it's something usual using golang.

I thought about having a json files containing the statistics and achievements rules,- as I already saw something like that in go projects made by Google.

So that's how I introduced the file `rules.json` where we will find facts about statistics and achievements.

```
{
    "statistics": {
        "game": [
			...
        ],
        "historical": [
			...
        ]
    },
    "achievements": {
        "rules": [
			...
        ]
    }
}
```

In my opinion, it makes sense to have a system working like that, it implies that we have to do a new deployment of the achievements system when introducing a new set of rules (stats/achievements), but at least we can track them with the versionning system.
And also, we can use the fact that Golang is strongly typed.