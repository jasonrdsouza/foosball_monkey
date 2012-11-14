package main

import (
    "fmt"
    "net/http"
    "strconv"
    "html/template"
    "encoding/json"
    "github.com/jasonrdsouza/foosball_monkey/datastore"
    "github.com/gorilla/mux"
    "github.com/gorilla/schema"
    //"time"
)

// Change this to use different underlying datastore
var database = datastore.Sqlite3DataHandler{}
var db = datastore.FoosballMonkeyDataHandler(database)

// html form decoder
var decoder = schema.NewDecoder()

var index_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/index.html",
))
var players_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/players.html",
))
var addplayer_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/add_player.html",
))
var deleteplayer_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/delete_player.html",
))
var player_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/player.html",
))
var games_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/games.html",
))
var addgame_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/add_game.html",
))
var deletegame_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/delete_game.html",
))
var game_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/game.html",
))
var teams_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/teams.html",
))
var addteam_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/add_team.html",
))
var deleteteam_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/delete_team.html",
))

var temp_temp, err = template.ParseFiles(
    "templates/_base.html",
    "templates/team.html",
)

var team_html = template.Must(template.New("team_template").Funcs(
        template.FuncMap{
            "mod0": func(a, b int) bool {
                return a % b == 0
            },
        }).ParseFiles(
    "templates/_base.html",
    "templates/team.html",
))

var queue_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/queue.html",
))
var rankings_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/rankings.html",
))


func home(w http.ResponseWriter, req *http.Request) {
    if err := index_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func connectToDB() {
    err := db.ConnectToDB("foosball_monkey_datastore.db")
}

func addPlayerHandler(w http.ResponseWriter, req *http.Request) {
    team_id, _ := strconv.Atoi(req.FormValue("team"))
    err := datastore.AddPlayer(database, req.FormValue("name"), req.FormValue("email"), req.FormValue("tagline"), team_id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/players", http.StatusFound)
}

func addPlayer(w http.ResponseWriter, req *http.Request) {
    //get teams to populate the form
    var teams []datastore.Team
    teams, err := datastore.GetAllTeams(database)
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
    players, err := datastore.GetAllPlayers(database)
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
    players, err := datastore.GetAllPlayers_display(database)
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
    p, err := datastore.GetPlayerByID(database, id)
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
    p, err := datastore.GetPlayerByID(database, id)
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
    players, err := datastore.GetAllPlayers(database)
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

    err := datastore.DeletePlayer(database, player_id)
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

    err := datastore.AddGame(database, offenderA, 
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
    players, err := datastore.GetAllPlayers(database)
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
    var games []datastore.GameDisplay
    games, err := datastore.GetAllGames_display(database)
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
    var games []datastore.Game
    games, err := datastore.GetAllGames(database)
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
    g, err := datastore.GetGameByID(database, id)
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
    g, err := datastore.GetGameByID(database, id)
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
    var games []datastore.GameDisplay
    games, err := datastore.GetAllGames_display(database)
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

    err := datastore.DeleteGame(database, game_id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/games", http.StatusFound)
}

func getAllTeams(w http.ResponseWriter, req *http.Request) {
    var teams []datastore.Team
    teams, err := datastore.GetAllTeams(database)
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
    teams, err := datastore.GetAllTeams(database)
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

    err := datastore.AddTeam(database, name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/teams", http.StatusFound)
}

func getTeamByIdJSON(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    t, err := datastore.GetTeamByID(database, id)
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
    t, err := datastore.GetTeamByID(database, id)
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
    teams, err := datastore.GetAllTeams(database)
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

    err := datastore.DeleteTeam(database, team_id)
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


func main() {
    //initializeDB()  //uncomment this to remake the database
    connectToDB()

    r := mux.NewRouter()

    r.HandleFunc("/", home)
    r.HandleFunc("/players", getAllPlayers)
    r.HandleFunc("/players.json", getAllPlayersJSON)
    r.HandleFunc("/players/{id:[0-9]+}.json", getPlayerByIdJSON)
    r.HandleFunc("/players/{id:[0-9]+}", getPlayerById)
    r.HandleFunc("/players/add", addPlayer)
    r.HandleFunc("/players/addHandler", addPlayerHandler)
    r.HandleFunc("/players/delete", deletePlayer)
    r.HandleFunc("/players/deleteHandler", deletePlayerHandler)
    r.HandleFunc("/games", getAllGames)
    r.HandleFunc("/games.json", getAllGamesJSON)
    r.HandleFunc("/games/{id:[0-9]+}.json", getGameByIdJSON)
    r.HandleFunc("/games/{id:[0-9]+}", getGameById)
    r.HandleFunc("/games/add", addGame)
    r.HandleFunc("/games/addHandler", addGameHandler)
    r.HandleFunc("/games/delete", deleteGame)
    r.HandleFunc("/games/deleteHandler", deleteGameHandler)
    r.HandleFunc("/teams", getAllTeams)
    r.HandleFunc("/teams.json", getAllTeamsJSON)
    r.HandleFunc("/teams/{id:[0-9]+}.json", getTeamByIdJSON)
    r.HandleFunc("/teams/{id:[0-9]+}", getTeamById)
    r.HandleFunc("/teams/add", addTeam)
    r.HandleFunc("/teams/addHandler", addTeamHandler)
    r.HandleFunc("/teams/delete", deleteTeam)
    r.HandleFunc("/teams/deleteHandler", deleteTeamHandler)
    r.HandleFunc("/queue", getQueue)
    r.HandleFunc("/queue.json", getQueueJSON)
    r.HandleFunc("/rankings", getRankings)
    r.HandleFunc("/rankings.json", getRankingsJSON)
    r.HandleFunc("/search/{query:[a-z0-9]+}.json", getSearchResultsJSON)

    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

    http.Handle("/", r)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}