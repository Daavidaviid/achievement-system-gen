# Achievements system for a fictional online game
# David Guerin's solution

## Explanations

The goal of this project was to make an extensible achievements system. I decided that I would do it using code generation, as it allows us to make an extensible system while keeping it strongly typed. That way, when one wants to add a new statistics (be it either for a game or historical) or new achievement, one needs to update the `rules.json` file accordingly.

It needs to be done in conjunction with the team working on the game, so that they can produce the new statistic necessary for the new achievement.

## Packages
Be careful, you need to run the program at least once to see the all generated files.
### achievementManager
This package is the goal of this project and is it used to manage the users achievements after a game. It can display them in the console or return json data.
### database
This package role is to mimicate a real database where we would store users data and historical players statistics. We could also store each game state once they are done.
### driver
The role of this package is to show that the achievements system works. It creates dummy data, loads it in the "database", create a dummy Game with people from the database, and then send the game data to the achievements system once the game is "done" .
### gen
The role of this package is to generate go sources according to the `rules.json`.
### gen_files
This package contains functions used by the `gen` package, noticeably there is the function `CheckRulesValidity(rules *Rules) bool` to make sure the rules defined in `rules.json` are legit.
### models
In this package we can find the models necessary for this project.

## Getting started

To launch the program, run :

```
./run.sh
```

This program will do the following things in this very specific order :

1. Generates the necessary `*_gen.go` files using the rules defined in `rules.json` if the rules are valid.
2. Run the tests accordingly.
3. Create a system with dummy users and their respective player historical statistics.
4. Pick some users to create a game.
5. Develop the game by increasing the score and each player statistics.
6. End the game and send the game data to the achievements system.
7. The achievements system display the achievements for each player (game and historical):  
    1. all achievements
    1. Freshly unlocked achievements
8. The achievements system update the database accordingly.

## How to add a new statistic / achievement

Into the `rules.json` file, you can add either a game `statistics` or a `historical` statistics.

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

To add a new statistic, go into the `achievements.statistics` section, For example, to add a `numberOfJump` statistic, we can add this : 
```
{
    "name": "NumberOfJump",
    "key": "numberOfJump",
    "type": "int",
    "description": "Number of jumps"
}
```

To add a new achievement, go into the `achievements.rules` section, For example, to add a `Jumpy Killer` achievement, we can add this : 

```
{
    "key": "JumpyKiller",
    "name": "Jumpy Killer",
    "description": "A user receives this for killing more than 30 enemies during a game and jumping at least 300 in one game",
    "rule": "game.NumberOfKills >= 30 && game.NumberOfJump >= 300",
    "tests": {
        "success": {
            "game": {
                "NumberOfKills": 30,
                "NumberOfJump": 450
            }
        },
        "failure": {
            "game": {
                "NumberOfKills": 10,
                "NumberOfJump": 450
            }
        }
    }
}
```