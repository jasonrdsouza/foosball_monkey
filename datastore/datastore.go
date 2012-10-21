package datastore

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "os"
    "time"
    "strings"
    "strconv"
    "errors"
    "crypto/md5"
    "io"
    "fmt"
)


type Player struct {
    Id int
    Name string
    Email string
    Email_md5 string
    Tagline string
    Team int
}

type Game struct {
    Id int
    OffenderA int
    DefenderA int
    OffenderB int
    DefenderB int
    ScoreA int
    ScoreB int
    Winner string
    Timestamp time.Time
}

type GameDisplay struct {
    Id int
    OffenderA string
    DefenderA string
    OffenderB string
    DefenderB string
    ScoreA int
    ScoreB int
    Winner string
    Timestamp time.Time
}

type Team struct {
    Id int
    Name string
    Members []Player
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
        "create table players (id integer not null primary key, name text, email text, email_md5 text, tagline text, team int)",
        "create table games (id integer not null primary key, offenderA integer, defenderA integer, offenderB integer, defenderB integer, scoreA integer, scoreB integer, winner string, dt datetime)",
        "create table teams (id integer not null primary key, name text)",
    }
    for _, sql := range sqls {
        _, err = db.Exec(sql)
        if err != nil {
            return err
        }
    }
    return nil
}

func BackupDB(db_name string) error {
    //since this is sqlite, we can simply copy the db file to backup
    return errors.New("Backing up the database is not currently supported")
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

func AddPlayer(db_name string, player_name string, email string, tagline string, team int) (error) {
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
    stmt, err := tx.Prepare("insert into players(name, email, email_md5, tagline, team) values(?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    
    //generate email md5 value
    h := md5.New()
    io.WriteString(h, email)
    email_md5 := fmt.Sprintf("%x", h.Sum(nil))
    
    _, err = stmt.Exec(player_name, email, email_md5, tagline, team)
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

    rows, err := db.Query("select id, name, email, email_md5, tagline, team from players")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    players := make([]Player, 0)
    for rows.Next() {
        player := Player{}
        rows.Scan(&(player.Id), &(player.Name), &(player.Email), &(player.Email_md5), &(player.Tagline), &(player.Team))
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

    stmt, err := db.Prepare("select name, email, email_md5, tagline, team from players where id = ?")
    if err != nil {
        return Player{}, err
    }
    defer stmt.Close()

    var name, email, email_md5, tagline string
    var team int
    err = stmt.QueryRow(id).Scan(&name, &email, &email_md5, &tagline, &team)
    if err != nil {
        return Player{}, err 
    }

    fetched_player := Player{id, name, email, email_md5, tagline, team}
    return fetched_player, nil
}

func AddGame(db_name string, offenderA int, defenderA int, offenderB int, defenderB int, scoreA int, scoreB int, winner string, dt string) (error) {
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
    stmt, err := tx.Prepare("insert into games(offenderA, defenderA, offenderB, defenderB, scoreA, scoreB, winner, dt) values(?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(offenderA, defenderA, offenderB, defenderB, scoreA, scoreB, winner, dt)
    if err != nil {
        return err 
    }
    tx.Commit()

    return nil
}

func GetAllGames(db_name string) ([]Game, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("select id, offenderA, defenderA, offenderB, defenderB, scoreA, scoreB, winner, dt from games")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    games := make([]Game, 0)
    for rows.Next() {
        game := Game{}
        var date_string string
        rows.Scan(&(game.Id), &(game.OffenderA), &(game.DefenderA), &(game.OffenderB), &(game.DefenderB), &(game.ScoreA), &(game.ScoreB), &(game.Winner), &(date_string))
        game.Timestamp, _ = convertDateStrToTime(date_string)
        games = append(games, game)
    }

    return games, nil
}

func GetAllGames_display(db_name string) ([]GameDisplay, error) {
    games, err := GetAllGames(db_name)
    if err != nil {
        return nil, err
    }
    games_display := make([]GameDisplay, 0)
    for _, game := range games {
        offA, err := GetPlayerByID(db_name, game.OffenderA)
        if err != nil {
            return nil, err
        }
        defA, err := GetPlayerByID(db_name, game.DefenderA)
        if err != nil {
            return nil, err
        }
        offB, err := GetPlayerByID(db_name, game.OffenderB)
        if err != nil {
            return nil, err
        }
        defB, err := GetPlayerByID(db_name, game.DefenderB)
        if err != nil {
            return nil, err
        }
        gd := GameDisplay{game.Id, offA.Name, defA.Name, offB.Name, defB.Name, game.ScoreA, game.ScoreB, game.Winner, game.Timestamp}
        games_display = append(games_display, gd)
    }
    return games_display, nil
}

func GetGameByID(db_name string, id int) (Game, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return Game{}, err
    }
    defer db.Close()

    stmt, err := db.Prepare("select offenderA, defenderA, offenderB, defenderB, scoreA, scoreB, winner, dt from games where id = ?")
    if err != nil {
        return Game{}, err
    }
    defer stmt.Close()

    game := Game{}
    game.Id = id
    var date_string string
    err = stmt.QueryRow(id).Scan(&(game.OffenderA), &(game.DefenderA), &(game.OffenderB), &(game.DefenderB), &(game.ScoreA), &(game.ScoreB), &(game.Winner), &(date_string))
    if err != nil {
        return Game{}, err 
    }
    game.Timestamp, _ = convertDateStrToTime(date_string)

    return game, nil
}

func GetAllTeams(db_name string) ([]Team, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("select id, name from teams")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    teams := make([]Team, 0)
    for rows.Next() {
        team := Team{}
        rows.Scan(&(team.Id), &(team.Name))
        team.Members, _ = GetTeamMembers(db_name, team)
        teams = append(teams, team)
    }

    return teams, nil
}

func GetTeamByID(db_name string, id int) (Team, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return Team{}, err
    }
    defer db.Close()

    stmt, err := db.Prepare("select name from teams where id = ?")
    if err != nil {
        return Team{}, err
    }
    defer stmt.Close()

    team := Team{}
    team.Id = id
    err = stmt.QueryRow(id).Scan(&(team.Name))
    if err != nil {
        return Team{}, err 
    }

    players, err := GetTeamMembers(db_name, team)
    if err != nil {
        return team, err 
    }

    team.Members = players

    return team, nil
}

func GetTeamMembers(db_name string, team Team) ([]Player, error) {
    db_name = "./" + db_name
    db, err := sql.Open("sqlite3", db_name)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    stmt, err := db.Prepare("select id, name, email, email_md5, tagline, team from players where team = ?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    players := make([]Player, 0)
    rows, err := stmt.Query(team.Id)
    if err != nil {
        return nil, err
    }

    for rows.Next() {
        player := Player{}
        rows.Scan(&(player.Id), &(player.Name), &(player.Email), &(player.Email_md5), &(player.Tagline), &(player.Team))
        players = append(players, player)
    }

    return players, nil
}

func AddTeam(db_name string, team_name string) (error) {
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
    stmt, err := tx.Prepare("insert into teams(name) values(?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(team_name)
    if err != nil {
        return err 
    }
    tx.Commit()

    return nil
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