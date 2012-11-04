foosball_monkey
===============

Webapp to track foosball things.

Todo (for v1.0)
---------------
- backup db functionality
    - copy file somewhere else (dropbox?)
- GetGameById html style
    - pictures of the players
    - seperated by team
    - schematic of the foosball table?
- Implement basic rankings
- Remove the delete player and delete team functionality
- consolidate player and player_display struct into 1 that has all the info
    - same for others
    - need foreign keys as well as the values for links


Future Functionality
--------------------
- Integrate with third party services
    - similar to how I did with gravatar
- Change the redirect after adding a player/game/team
    - redirect to a page that acknowledge recieving the input
        - show error if duplicate or something
    - use a bootstrap banner thing?
- Get site search to work
- Fix stupid header formatting thing
    - that causes the bar at the top to overlap with the header of a page
- Better date validation
- Implement the queue
- Ability to email things
    - weekly report?
        - how many games you have played
        - your current ranking
        - etc
    - email errors/ logs to me
    - email feedback from the users to me
- Spruce up index page
- Don't allow deletion of teams/ players if there are references to them elsewhere?
- GetTeamById html style
    - gravatars of all the players


Ranking
-------
- generating rankings
    - every so often, kick off a process that calculates rankings
        - goes through all the games played
        - generates a rank for all the players/ teams
        - updates # of games played for all players/ teams
    - potential incremental update option
- algo for rank that takes amount of games played into account
    - elo type ranking
    - rank teams just like players
- Player rankings
    - list of top ranking players
    - # of games played by players (bar chart)
        - top 10?
    - player rank over time
- Team rankings
    - list of top ranking teams
    - breakdown of #players in each team (bar chart)
    - breakdown of # of games played by team
    - team rank over time (normalized line chart)
        - one line for each team
- Add ranking metrics to individual team/ player html pages
    - show current rank on the player or teams page
    - show number of games played on the player or teams page
    - show number of players on the team page
