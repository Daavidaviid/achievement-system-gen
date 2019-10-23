# Backend Coding Test

Thank you for taking the time to do our technical test.
Please make this a zip file containing one folder containing the files.
Coding Test
For your coding challenge, you will create an achievements system for a fictional online game.
Our fictional game is played online between two teams. Each team can have 3-5 players (but must be evenly matched – i.e., if one team has 4 players, the other must have 4 players).
Statistics are compiled for each player during/after each game:

1. Number of attempted attacks
2. Number of hits
3. Total Amount of damage done
4. Number of kills
5. Number of “first hit” kills
6. Number of assists
7. Total number of spells cast
8. Total spell damage done
9. Total time played

In addition, historical player statistics are maintained as well. For instance, each player is credited with a win/loss based upon the performance of his/her team during the game. Other historical player-level stats can be tracked as well – such as total number of games played, total duration of games played, total number of kills, etc. Achievements are awarded based upon any combination of these statistics (historical OR gamespecific). Players collect achievements over time, and these will be displayed on their profile page (the display is outside the scope of this exercise). This profile page is visible to all other players, so achievements act as a kind of “badge of honor”. After
 each game has ended, the achievements logic is called and achievements are awarded to each of the players if applicable.
Your mission, if you choose to accept it, is to create this achievements system. So, here’s what’s expected from this coding exercise:

1. Code should be written in Go, Java, C#, or C++.
2. The achievements system that you are creating should correctly assign achievements to each player involved in the game that just ended, if applicable.
3. Persistence, messaging, and other potential concerns are considered ancillary to this exercise, so don’t worry about implementing those.
4. Create all domain objects suggested in the above description, using the proper data structures where applicable.
5. Create the achievements system. At a minimum, it should have an entry point that would be called at the end of a game. You must include the following achievements:

a. “Sharpshooter Award” – a user receives this for landing 75% of their attacks, assuming they have at least attacked once.

b. “Bruiser” Award – a user receives this for doing more than 500 points of damage in one game

c. “Veteran” Award – a user receives this for playing more than 1000 games in their lifetime.

d. “Big Winner” Award – a user receives this for having 200 wins

6. Your system should be extensible in the following ways:

a. The system must easily handle the tracking of additional statistics at any level listed above (historical by player or game-specific by player)

b. The system must easily handle the addition of new achievements

7. Add a new statistic, and create a new Achievement utilizing this statistic in conjunction to one of the other statistics defined above.

8. Create unit tests to demonstrate that your code is functional using (e.g. using JUnit or equivalent for C#/C++)

9. Create a driver class that sets up sample data and calls the achievements system, printing out achievements to the console. Don’t worry about persistence, and use any method of creating/loading sample data that you are familiar with. Also, you can combine/integrate this into your unit tests if you prefer – we’re just looking for you to demonstrate that the code runs.

## Please take your time and treat this challenge as if it were for a professional project for a client; we're looking for more than simply working code.
## Thanks for your time, we look forward to hearing from you!