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

func home(w http.ResponseWriter, req *http.Request) {
    if err := index_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func initializeDB() {
    err := datastore.CreateNewDB(database)
    if err != nil {
        fmt.Println(err)
    }
}

func addPlayerHandler(w http.ResponseWriter, req *http.Request) {
    err := datastore.AddPlayer(database, req.FormValue("Name"), req.FormValue("Email"), req.FormValue("Tagline"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    http.Redirect(w, req, "/players/add", http.StatusCreated)
}

func addPlayer(w http.ResponseWriter, req *http.Request) {
    if err := addplayer_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getAllPlayersJSON(w http.ResponseWriter, req *http.Request) {
    players, err := datastore.GetAllPlayers(database)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    players_json, err := json.Marshal(players)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    fmt.Fprintln(w, string(players_json))
}

func getAllPlayers(w http.ResponseWriter, req *http.Request) {
    players, err := datastore.GetAllPlayers(database)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    if err := players_html.Execute(w, players); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getPlayerByIdJSON(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    p, err := datastore.GetPlayerByID(database, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    player_json, err := json.Marshal(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    fmt.Fprintln(w, string(player_json))
}

func addGameHandler(w http.ResponseWriter, req *http.Request) {
    offenderA, _ := strconv.Atoi(req.FormValue("OffenderA"))
    defenderA, _ := strconv.Atoi(req.FormValue("DefenderA"))
    offenderB, _ := strconv.Atoi(req.FormValue("OffenderB"))
    defenderB, _ := strconv.Atoi(req.FormValue("DefenderB"))
    scoreA, _ := strconv.Atoi(req.FormValue("ScoreA"))
    scoreB, _ := strconv.Atoi(req.FormValue("ScoreB"))
    winner := req.FormValue("Winner")
    timestamp := req.FormValue("Timestamp")

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
    }
    http.Redirect(w, req, "/games/add", http.StatusFound)
}

func addGame(w http.ResponseWriter, req *http.Request) {
    if err := addgame_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getAllGames(w http.ResponseWriter, req *http.Request) {
    var games []datastore.Game
    games, err := datastore.GetAllGames(database)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    if err := games_html.Execute(w, games); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getAllGamesJSON(w http.ResponseWriter, req *http.Request) {
    var games []datastore.Game
    games, err := datastore.GetAllGames(database)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    
    games_json, err := json.Marshal(games)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    fmt.Fprintln(w, string(games_json))
}

func getGameByIdJSON(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    g, err := datastore.GetGameByID(database, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    
    game_json, err := json.Marshal(g)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    fmt.Fprintln(w, string(game_json))
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
    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

    http.Handle("/", r)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}