package main

import (
    "fmt"
    "net/http"
    "strconv"
    "html/template"
    "encoding/json"
    "github.com/jasonrdsouza/foosball_monkey/datastore"
    "code.google.com/p/gorilla/mux"
    "code.google.com/p/gorilla/schema"
    //"time"
)

const database = "foosball_monkey_datastore.db"

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
var games_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/games.html",
))
var addgame_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/add_game.html",
))
var teams_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/teams.html",
))
var addteam_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/add_team.html",
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

func initializeDB() {
    err := datastore.CreateNewDB(database)
    if err != nil {
        fmt.Println(err)
        return
    }
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
    if err := addplayer_html.Execute(w, nil); err != nil {
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
    players, err := datastore.GetAllPlayers(database)
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
    if err := addgame_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

func getAllGames(w http.ResponseWriter, req *http.Request) {
    var games []datastore.Game
    games, err := datastore.GetAllGames(database)
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



func main() {
    //initializeDB()  //uncomment this to remake the database

    r := mux.NewRouter()

    r.HandleFunc("/", home)
    r.HandleFunc("/players", getAllPlayers)
    r.HandleFunc("/players.json", getAllPlayersJSON)
    r.HandleFunc("/players/{id:[0-9]+}.json", getPlayerByIdJSON)
    r.HandleFunc("/players/add", addPlayer)
    r.HandleFunc("/players/addHandler", addPlayerHandler)
    r.HandleFunc("/games", getAllGames)
    r.HandleFunc("/games.json", getAllGamesJSON)
    r.HandleFunc("/games/{id:[0-9]+}.json", getGameByIdJSON)
    r.HandleFunc("/games/add", addGame)
    r.HandleFunc("/games/addHandler", addGameHandler)
    r.HandleFunc("/teams", getAllTeams)
    r.HandleFunc("/teams.json", getAllTeamsJSON)
    r.HandleFunc("/teams/{id:[0-9]+}.json", getTeamByIdJSON)
    r.HandleFunc("/teams/add", addTeam)
    r.HandleFunc("/teams/addHandler", addTeamHandler)
    r.HandleFunc("/queue", getQueue)
    r.HandleFunc("/queue.json", getQueueJSON)
    r.HandleFunc("/rankings", getRankings)
    r.HandleFunc("/rankings.json", getRankingsJSON)

    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

    http.Handle("/", r)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}