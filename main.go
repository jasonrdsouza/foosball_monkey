package main

import (
    "fmt"
    "net/http"
    "strconv"
    "github.com/jasonrdsouza/foosball_monkey/datastore"
    "code.google.com/p/gorilla/mux"
    //"time"
)

const database = "foosball_monkey_datastore.db"

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello foosball world!")
}

func initializeDB() {
    err := datastore.CreateNewDB(database)
    if err != nil {
        fmt.Println(err)
    }
}

func addPlayer(w http.ResponseWriter, req *http.Request) {
    err := datastore.AddPlayer(database, "testplayer", "testing the tagline!")
    if err != nil {
        fmt.Fprintln(w, err)
    }
    fmt.Fprintln(w, "Test player added!")
}

func getAllPlayers(w http.ResponseWriter, req *http.Request) {
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

func addGame(w http.ResponseWriter, req *http.Request) {
    err := datastore.AddGame(database, 1, 1, 2, 2, 5, 10, "a", "2012-09-26")
    if err != nil {
        fmt.Fprintln(w, err)
    }
    fmt.Fprintln(w, "Test game added!")
}

func getAllGames(w http.ResponseWriter, req *http.Request) {
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
    initializeDB()
    r := mux.NewRouter()
    r.HandleFunc("/", hello)
    r.HandleFunc("/players", getAllPlayers)
    r.HandleFunc("/players/{id:[0-9]+}", getPlayerByID)
    r.HandleFunc("/players/add", addPlayer)
    r.HandleFunc("/games", getAllGames)
    r.HandleFunc("/games/{id:[0-9]+}", getGameByID)
    r.HandleFunc("/games/add", addGame)

    http.Handle("/", r)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}