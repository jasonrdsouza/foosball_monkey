package main

import (
  "fmt"
  "encoding/json"
  "strconv"
  "github.com/gorilla/mux"
  "net/http"
  "github.com/jasonrdsouza/foosball_monkey/datastore"
)


func home(w http.ResponseWriter, req *http.Request) {
    if err := index_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func addPlayerHandler(w http.ResponseWriter, req *http.Request) {
    team_id, _ := strconv.Atoi(req.FormValue("team"))
    err := db.AddPlayer(req.FormValue("name"), req.FormValue("email"), req.FormValue("tagline"), team_id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/players", http.StatusFound)
}

func addPlayer(w http.ResponseWriter, req *http.Request) {
    //get teams to populate the form
    var teams []datastore.Team
    teams, err := db.GetAllTeams()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := addplayer_html.Execute(w, teams); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func getAllPlayersJSON(w http.ResponseWriter, req *http.Request) {
    players, err := db.GetAllPlayers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    players_json, err := json.Marshal(players)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, string(players_json))
}

func getAllPlayers(w http.ResponseWriter, req *http.Request) {
    players, err := db.GetAllPlayers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := players_html.Execute(w, players); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func getPlayerByIdJSON(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    p, err := db.GetPlayerByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    player_json, err := json.Marshal(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, string(player_json))
}

func getPlayerById(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    p, err := db.GetPlayerByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := player_html.Execute(w, p); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func deletePlayer(w http.ResponseWriter, req *http.Request) {
    var players []datastore.Player
    players, err := db.GetAllPlayers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := deleteplayer_html.Execute(w, players); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func deletePlayerHandler(w http.ResponseWriter, req *http.Request) {
    player_id, _ := strconv.Atoi(req.FormValue("name"))

    err := db.DeletePlayer(player_id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/players", http.StatusFound)
}

func addGameHandler(w http.ResponseWriter, req *http.Request) {
    offenderA, _ := strconv.Atoi(req.FormValue("offenderA"))
    defenderA, _ := strconv.Atoi(req.FormValue("defenderA"))
    offenderB, _ := strconv.Atoi(req.FormValue("offenderB"))
    defenderB, _ := strconv.Atoi(req.FormValue("defenderB"))
    scoreA, _ := strconv.Atoi(req.FormValue("scoreA"))
    scoreB, _ := strconv.Atoi(req.FormValue("scoreB"))
    winner := req.FormValue("winner")
    timestamp := req.FormValue("timestamp")

    err := db.AddGame(offenderA, 
                       defenderA, 
                       offenderB, 
                       defenderB, 
                       scoreA, 
                       scoreB, 
                       winner, 
                       timestamp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/games", http.StatusFound)
}

func addGame(w http.ResponseWriter, req *http.Request) {
    //get players for dropdown
    players, err := db.GetAllPlayers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := addgame_html.Execute(w, players); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func getAllGames(w http.ResponseWriter, req *http.Request) {
    games, err := db.GetAllGames()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := games_html.Execute(w, games); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func getAllGamesJSON(w http.ResponseWriter, req *http.Request) {
    games, err := db.GetAllGames()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    games_json, err := json.Marshal(games)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, string(games_json))
}

func getGameByIdJSON(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    g, err := db.GetGameByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    game_json, err := json.Marshal(g)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, string(game_json))
}

func getGameById(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    g, err := db.GetGameByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := game_html.Execute(w, g); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func deleteGame(w http.ResponseWriter, req *http.Request) {
    var games []datastore.Game
    games, err := db.GetAllGames()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := deletegame_html.Execute(w, games); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func deleteGameHandler(w http.ResponseWriter, req *http.Request) {
    game_id, _ := strconv.Atoi(req.FormValue("name"))

    err := db.DeleteGame(game_id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/games", http.StatusFound)
}

func getAllTeams(w http.ResponseWriter, req *http.Request) {
    var teams []datastore.Team
    teams, err := db.GetAllTeams()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := teams_html.Execute(w, teams); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func getAllTeamsJSON(w http.ResponseWriter, req *http.Request) {
    var teams []datastore.Team
    teams, err := db.GetAllTeams()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    teams_json, err := json.Marshal(teams)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, string(teams_json))
}

func addTeam(w http.ResponseWriter, req *http.Request) {
    if err := addteam_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func addTeamHandler(w http.ResponseWriter, req *http.Request) {
    name := req.FormValue("name")

    err := db.AddTeam(name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/teams", http.StatusFound)
}

func getTeamByIdJSON(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    t, err := db.GetTeamByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    team_json, err := json.Marshal(t)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, string(team_json))
}

func getTeamById(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    t, err := db.GetTeamByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := team_html.Execute(w, t); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func deleteTeam(w http.ResponseWriter, req *http.Request) {
    var teams []datastore.Team
    teams, err := db.GetAllTeams()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := deleteteam_html.Execute(w, teams); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func deleteTeamHandler(w http.ResponseWriter, req *http.Request) {
    team_id, _ := strconv.Atoi(req.FormValue("name"))

    err := db.DeleteTeam(team_id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/teams", http.StatusFound)
}

func getQueue(w http.ResponseWriter, req *http.Request) {
    if err := queue_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func getQueueJSON(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Implement this JSON queue functionality")
}

func getRankings(w http.ResponseWriter, req *http.Request) {
    if err := rankings_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func getRankingsJSON(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Implement this JSON ranking functionality")
}

func getSearchResultsJSON(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    query, _ := vars["query"]
    // do some database stuff here
    
    temp_output := fmt.Sprintf("You searched for: %s", query)
    temp_output_json, err := json.Marshal(temp_output)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, string(temp_output_json))
}

func createRouter() (*mux.Router, error) {
    r := mux.NewRouter()

    r.HandleFunc("/", home)
    r.HandleFunc("/players", getAllPlayers)
    r.HandleFunc("/players.json", getAllPlayersJSON)
    r.HandleFunc("/players/{id:[0-9]+}.json", getPlayerByIdJSON)
    r.HandleFunc("/players/{id:[0-9]+}", getPlayerById)
    r.HandleFunc("/players/add", addPlayer)
    r.HandleFunc("/players/addHandler", addPlayerHandler)
    //r.HandleFunc("/players/delete", deletePlayer)
    //r.HandleFunc("/players/deleteHandler", deletePlayerHandler)
    r.HandleFunc("/games", getAllGames)
    r.HandleFunc("/games.json", getAllGamesJSON)
    r.HandleFunc("/games/{id:[0-9]+}.json", getGameByIdJSON)
    r.HandleFunc("/games/{id:[0-9]+}", getGameById)
    r.HandleFunc("/games/add", addGame)
    r.HandleFunc("/games/addHandler", addGameHandler)
    //r.HandleFunc("/games/delete", deleteGame)
    //r.HandleFunc("/games/deleteHandler", deleteGameHandler)
    r.HandleFunc("/teams", getAllTeams)
    r.HandleFunc("/teams.json", getAllTeamsJSON)
    r.HandleFunc("/teams/{id:[0-9]+}.json", getTeamByIdJSON)
    r.HandleFunc("/teams/{id:[0-9]+}", getTeamById)
    r.HandleFunc("/teams/add", addTeam)
    r.HandleFunc("/teams/addHandler", addTeamHandler)
    //r.HandleFunc("/teams/delete", deleteTeam)
    //r.HandleFunc("/teams/deleteHandler", deleteTeamHandler)
    r.HandleFunc("/queue", getQueue)
    r.HandleFunc("/queue.json", getQueueJSON)
    r.HandleFunc("/rankings", getRankings)
    r.HandleFunc("/rankings.json", getRankingsJSON)
    r.HandleFunc("/search/{query:[a-z0-9]+}.json", getSearchResultsJSON)

    return r, nil
}