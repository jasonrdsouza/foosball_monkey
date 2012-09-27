package datastore

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "os"
    "time"
    "strings"
    "strconv"
    "errors"
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
    PlayerA int
    PlayerB int
    ScoreA int
    ScoreB int
    Winner string
    Timestamp time.Time
}

type Game1v2 struct {
    Id int
    PlayerA int
    PlayerB1 int
    PlayerB2 int
    ScoreA int
    ScoreB int
    Winner string
    Timestamp time.Time
}

type Game2v2 struct {
    Id int
    PlayerA1 int
    PlayerA2 int
    PlayerB1 int
    PlayerB2 int
    ScoreA int
    ScoreB int
    Winner string
    Timestamp time.Time
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
        "create table games1v1 (id integer not null primary key, PlayerA integer, PlayerB integer, ScoreA integer, ScoreB integer, winner string, dt datetime)",
        "create table games1v2 (id, integer not null primary key, PlayerA integer, PlayerB1 integer, PlayerB2 integer, ScoreA integer, ScoreB integer, winner string, dt datetime)",
        "create table games2v2 (id, integer not null primary key, PlayerA1 integer, PlayerA2 integer, PlayerB1 integer, PlayerB2 integer, ScoreA integer, ScoreB integer, winner string, dt datetime)",
    }
    for _, sql := range sqls {
        _, err = db.Exec(sql)
        if err != nil {
            return err
        }
    }
    return nil
}

func convertDateStrToTime(date_string string) (time.Time, error) {
    dt_parts := strings.Split(date_string, "-")
    if len(dt_parts) < 3 {
        return time.Time{}, errors.New("Malformed date string detected")
    }
    year, _ := strconv.Atoi(dt_parts[0])
    month, _ := strconv.Atoi(dt_parts[1])
    day, _ := strconv.Atoi(dt_parts[2])
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
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

func AddGame1v1(db_name string, playerA int, playerB int, scoreA int, scoreB int, winner string, dt string) (error) {
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
    stmt, err := tx.Prepare("insert into games1v1(PlayerA, PlayerB, ScoreA, ScoreB, winner, dt) values(?, ?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(playerA, playerB, scoreA, scoreB, winner, dt)
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

    rows, err := db.Query("select id, PlayerA, PlayerB, ScoreA, ScoreB, winner, dt from games1v1")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    games := make([]Game1v1, 0)
    for rows.Next() {
        game := Game1v1{}
        var date_string string
        rows.Scan(&(game.Id), &(game.PlayerA), &(game.PlayerB), &(game.ScoreA), &(game.ScoreB), &(game.Winner), &(date_string))
        game.Timestamp, _ = convertDateStrToTime(date_string)
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

    stmt, err := db.Prepare("select PlayerA, PlayerB, ScoreA, ScoreB, winner, dt from games1v1 where id = ?")
    if err != nil {
        return Game1v1{}, err
    }
    defer stmt.Close()

    game := Game1v1{}
    game.Id = id
    var date_string string
    err = stmt.QueryRow(id).Scan(&(game.PlayerA), &(game.PlayerB), &(game.ScoreA), &(game.ScoreB), &(game.Winner), &(date_string))
    if err != nil {
        return Game1v1{}, err 
    }
    game.Timestamp, _ = convertDateStrToTime(date_string)

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