package main

import (
  "fmt"
  "net/http"
  "github.com/jasonrdsouza/foosball_monkey/datastore"
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

func main() {
  initializeDB()
  http.HandleFunc("/", hello)
  http.HandleFunc("/add", addPlayer)
  http.HandleFunc("/getAll", getAllPlayers)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}