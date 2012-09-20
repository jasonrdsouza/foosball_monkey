package main

import (
  "fmt"
  "net/http"
  "strconv"
  "github.com/jasonrdsouza/foosball_monkey/datastore"
  "code.google.com/p/gorilla/mux"
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
  err := datastore.AddPlayer(database, "testplayer")
  if err != nil {
    fmt.Fprintln(w, err)
  }
  fmt.Fprintln(w, "Test player added!")
}

func getAllPlayers(w http.ResponseWriter, req *http.Request) {
  player_string, err := datastore.GetAllPlayers(database)
  if err != nil {
    fmt.Fprintln(w, err)
  }
  fmt.Fprintln(w, player_string)
}

func getPlayerByID(w http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  id, _ := strconv.Atoi(vars["id"])
  fetched_player, err := datastore.GetPlayerByID(database, id)
  if err != nil {
    fmt.Fprintln(w, err)
  }
  fmt.Fprintln(w, fetched_player)
}

func addGame1v1(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "Implement addGame functionality!")
}

func getAllGames1v1(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "Implement getAllGames functionality!")
}

func getGame1v1ByID(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "Implement getGameByID functionality!")
}



func main() {
  initializeDB()
  r := mux.NewRouter()
  r.HandleFunc("/", hello)
  r.HandleFunc("/players", getAllPlayers)
  r.HandleFunc("/players/{id:[0-9]+}", getPlayerByID)
  r.HandleFunc("/players/add", addPlayer)
  r.HandleFunc("/games1v1", getAllGames1v1)
  r.HandleFunc("/games1v1/{id:[0-9]+}", getGame1v1ByID)
  r.HandleFunc("/games1v1/add", addGame1v1)

  http.Handle("/", r)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}