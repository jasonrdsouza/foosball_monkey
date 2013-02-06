// +build appengine

package main

import (
    "fmt"
    "net/http"
)

// Change this to use different underlying datastore
var database = datastore.AppengineDataHandler{}
var db = datastore.FoosballMonkeyDataHandler(&database)

func createNewDB(db_name string) {
    err := db.CreateNewDB(db_name)
    if err != nil {
        fmt.Println(err)
        return
    }
}

func connectToDB(db_name string) {
    err := db.ConnectToDB(db_name)
    if err != nil {
        fmt.Println(err)
        return
    }
}

func init() {
    //createNewDB("foosball_monkey_datastore.db")  //uncomment this to remake the database
    connectToDB("foosball_monkey_datastore.db")

    r, err := createRouter()
    if err != nil {
        panic(err)
    }

    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

    http.Handle("/", r)
}