foosball_monkey
===============

Webapp to track foosball things.

Todo (for v1.0)
---------------
- Add validations
    - Games
        - date/ time should be validated somehow
            - sync this with the go representation of date/time somehow
- backup db functionality
    - copy file somewhere else (dropbox?)
- Change the redirect after adding a player/game/team
    - redirect to a page that acknowledge recieving the input
        - show error if duplicate or something
    - use a bootstrap banner thing?
- GetPlayerById html style
- GetGameById html style
- GetTeamById html style
- Get site search to work
- Team members functionality
    - populate the teams table with the right members in the members column
- Implement delete functionality
    - for players
    - for games
    - for teams
    - for queue
- Implement the queue
- Implement the rankings
- Change the index page to reflect proper functionality
- Fix stupid header formatting thing
    - that causes the bar at the top to overlap with the header of a page


Future Functionality
--------------------
- Integrate with third party services
    - similar to how I did with gravatar


Ranking
-------
- algo for rank that takes amount of games played into account
- ranking by player and by team
- use a ranking system that takes into account the current rank of the opponents
- Player ranking
    - offensive vs defensive ranking
- Team ranking
    - how to rank teams?
