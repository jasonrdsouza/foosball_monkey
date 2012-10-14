foosball_monkey
===============

Webapp to track foosball things.

Todo
----
- Add game/ player form validations
    - change the game one to be a dropdown list of players
    - ensure the score is a number
    - winner dropdown (A or B)
    - Time input
    - use Bootstrap functionality
- backup db functionality
    - copy file somewhere else (dropbox?)
- Change the redirect after adding a player/game/team
    - redirect to a page that acknowledge recieving the input
    - use a bootstrap banner thing?
- GetPlayerById html style
- GetGameById html style
- GetTeamById html style
- Get site search to work
- Implement delete functionality
    - for players
    - for games
    - for teams
    - for queue
- Implement the queue
- Implement the rankings


Functionality
-------------
- Players
    - wins/ losses
    - ranking
        - within team
        - overall
        - defensive ranking?
        - offensive ranking
    - score of games played
    - number of games played
    - frequency of position played
        - offense vs defense?
- Integrate with third party services
    - similar to how I did with gravatar


Ranking
-------
- algo for rank that takes amount of games played into account
