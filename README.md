foosball_monkey
===============

Webapp to track foosball things.

Todo (for v1.0)
---------------
- backup db functionality
    - copy file somewhere else (dropbox?)
- GetPlayerById html style
- GetGameById html style
- GetTeamById html style
- Implement basic rankings


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


Ranking
-------
- algo for rank that takes amount of games played into account
- ranking by player and by team
- use a ranking system that takes into account the current rank of the opponents
- Player ranking
    - offensive vs defensive ranking
- Team ranking
    - how to rank teams?
