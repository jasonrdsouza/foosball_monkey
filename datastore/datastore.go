package datastore

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "os"
    //"strconv"
    //"errors"
    //"fmt"
)

/*
type Datastore interface {
    CreateNewDB
    AddPlayer
    Add1v1Game
}
*/

type Player struct {
    Id int
    Name string
}

type Game1v1 struct {
    Id int
    Player1 int
    Player2 int
    Score1 int
    Score2 int
    Winner int
}

func CreateNewDB(db_name string) error {
    db_name = "./" + db_name
    //remove old db
    os.Remove(db_name)

    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return err
    }
    defer db.Close()

    sqls := []string{
        "create table players (id integer not null primary key, name text)",
        "create table games1v1 (id integer not null primary key, player1 integer, player2 integer, score1 integer, score2 integer, winner integer)",
    }
    for _, sql := range sqls {
        _, err = db.Exec(sql)
        if err != nil {
            return err
        }
    }
    return nil
}

func AddPlayer(db_name string, player_name string) (error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return err
    }
    defer db.Close()

    //start a transaction (tx)
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    stmt, err := tx.Prepare("insert into players(name) values(?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(player_name)
    if err != nil {
        return err 
    }
    tx.Commit()
    return nil
}

func GetAllPlayers(db_name string) ([]Player, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("select id, name from players")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    players := make([]Player, 0)
    for rows.Next() {
        player := Player{}
        rows.Scan(&(player.Id), &(player.Name))
        players = append(players, player)
    }
    return players, nil
}

func GetPlayerByID(db_name string, id int) (Player, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return Player{}, err
    }
    defer db.Close()

    stmt, err := db.Prepare("select name from players where id = ?")
    if err != nil {
        return Player{}, err
    }
    defer stmt.Close()

    var name string
    err = stmt.QueryRow(id).Scan(&name)
    if err != nil {
        return Player{}, err 
    }

    fetched_player := Player{id, name}
    return fetched_player, nil
}

func AddGame1v1(db_name string, player1 int, player2 int, score1 int, score2 int, winner int) (error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return err
    }
    defer db.Close()

    //start a transaction (tx)
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    stmt, err := tx.Prepare("insert into games1v1(player1, player2, score1, score2, winner) values(?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(player1, player2, score1, score2, winner)
    if err != nil {
        return err 
    }
    tx.Commit()
    return nil
}

func GetAllGames1v1(db_name string) ([]Game1v1, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("select id, player1, player2, score1, score2, winner from games1v1")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    games := make([]Game1v1, 0)
    for rows.Next() {
        game := Game1v1{}
        rows.Scan(&(game.Id), &(game.Player1), &(game.Player2), &(game.Score1), &(game.Score2), &(game.Winner))
        games = append(games, game)
    }
    return games, nil
}

func GetGame1v1ByID(db_name string, id int) (Game1v1, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return Game1v1{}, err
    }
    defer db.Close()

    stmt, err := db.Prepare("select player1, player2, score1, score2, winner from games1v1 where id = ?")
    if err != nil {
        return Game1v1{}, err
    }
    defer stmt.Close()

    game := Game1v1{}
    game.Id = id
    err = stmt.QueryRow(id).Scan(&(game.Player1), &(game.Player2), &(game.Score1), &(game.Score2), &(game.Winner))
    if err != nil {
        return Game1v1{}, err 
    }

    return game, nil
}
    
/*    

    rows, err := db.Query("select id, name from foo")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()
    for rows.Next() {
        var id int
        var name string
        rows.Scan(&id, &name)
        println(id, name)
    }
    rows.Close()

    stmt, err = db.Prepare("select name from foo where id = ?")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()
    var name string
    err = stmt.QueryRow("3").Scan(&name)
    if err != nil {
        fmt.Println(err)
        return
    }
    println(name)

*/