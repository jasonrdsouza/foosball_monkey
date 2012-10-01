package main

import (
    "fmt"
    "net/http"
    "strconv"
    "html/template"
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
    err := datastore.AddPlayer(database, req.FormValue("Name"), req.FormValue("Tagline"))
    if err != nil {
        fmt.Fprintln(w, err)
    }
    http.Redirect(w, req, "/players/add", http.StatusCreated)
}

func addPlayer(w http.ResponseWriter, req *http.Request) {
    if err := addplayer_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getAllPlayersTXT(w http.ResponseWriter, req *http.Request) {
    players, err := datastore.GetAllPlayers(database)
    if err != nil {
        fmt.Fprintln(w, err)
    }
    output_string := ""
    for _, p := range players {
        output_string += fmt.Sprintf("Player %d: %s\n\tTagline: %s\n", p.Id, p.Name, p.Tagline)
    }
    fmt.Fprintln(w, output_string)
}

func getAllPlayers(w http.ResponseWriter, req *http.Request) {
    players, err := datastore.GetAllPlayers(database)
    if err != nil {
        fmt.Fprintln(w, err)
    }
    if err := players_html.Execute(w, players); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getPlayerByID(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    p, err := datastore.GetPlayerByID(database, id)
    if err != nil {
        fmt.Fprintln(w, err)
    }
    output_string := fmt.Sprintf("Player %d: %s\n\tTagline: %s\n", p.Id, p.Name, p.Tagline)
    fmt.Fprintln(w, output_string)
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
        fmt.Fprintln(w, err)
    }
    http.Redirect(w, req, "/games/add", http.StatusCreated)
}

func addGame(w http.ResponseWriter, req *http.Request) {
    if err := addgame_html.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getAllGamesTXT(w http.ResponseWriter, req *http.Request) {
    var games []datastore.Game
    games, err := datastore.GetAllGames(database)
    if err != nil {
        fmt.Fprintln(w, err)
    }
    
    output_string := ""
    for i, g := range games {
        output_string += fmt.Sprintf("Game %d:\n\tID: %d\n\tOffender A: %d\n\tDefender A: %d\n\tOffender B: %d\n\tDefender B: %d\n\tScore A: %d\n\tScore B: %d\n\tWinner: %s\n\tPlayed: %s\n", 
                                    i, g.Id, g.OffenderA, g.DefenderA, g.OffenderB, g.DefenderB, g.ScoreA, g.ScoreB, g.Winner, g.Timestamp)
    }
    fmt.Fprintln(w, output_string)
}

func getAllGames(w http.ResponseWriter, req *http.Request) {
    var games []datastore.Game
    games, err := datastore.GetAllGames(database)
    if err != nil {
        fmt.Fprintln(w, err)
    }
    if err := games_html.Execute(w, games); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getGameByID(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, _ := strconv.Atoi(vars["id"])
    g, err := datastore.GetGameByID(database, id)
    if err != nil {
        fmt.Fprintln(w, err)
    }
    output_string := fmt.Sprintf("Game ID: %d\n\tOffender A: %d\n\tDefender A: %d\n\tOffender B: %d\n\tDefender B: %d\n\tScore A: %d\n\tScore B: %d\n\tWinner: %s\n\tPlayed: %s\n", 
                                        g.Id, g.OffenderA, g.DefenderA, g.OffenderB, g.DefenderB, g.ScoreA, g.ScoreB, g.Winner, g.Timestamp)
    fmt.Fprintln(w, output_string)
}



func main() {
    //initializeDB()
    r := mux.NewRouter()
    r.HandleFunc("/", home)
    r.HandleFunc("/players", getAllPlayers)
    r.HandleFunc("/players/{id:[0-9]+}", getPlayerByID)
    r.HandleFunc("/players/add", addPlayer)
    r.HandleFunc("/players/addHandler", addPlayerHandler)
    r.HandleFunc("/games", getAllGames)
    r.HandleFunc("/games/{id:[0-9]+}", getGameByID)
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