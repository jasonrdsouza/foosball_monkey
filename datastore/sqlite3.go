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

type Sqlite3DataHandler struct {
    db_name string
    db *sql.DB
}

func (s Sqlite3DataHandler) CreateNewDB(db_name string) error {
    s.db_name = "./" + db_name
    //remove old db
    os.Remove(s.db_name)

    temp_db, err := sql.Open("sqlite3", s.db_name)
    s.db = temp_db
    if err != nil {
        return err
    }

    sqls := []string{
        "create table players (id integer not null primary key, name text, email text, email_md5 text, tagline text, team int)",
        "create table games (id integer not null primary key, offenderA integer, defenderA integer, offenderB integer, defenderB integer, scoreA integer, scoreB integer, winner string, dt datetime)",
        "create table teams (id integer not null primary key, name text)",
    }
    for _, sql := range sqls {
        _, err = (s.db).Exec(sql)
        if err != nil {
            return err
        }
    }
    return nil
}

func (s Sqlite3DataHandler) ConnectToDB(db_name string) error {
    s.db_name = "./" + db_name

    temp_db, err := sql.Open("sqlite3", s.db_name)
    s.db = temp_db
    if err != nil {
        return err
    }
    return nil
}

func (s Sqlite3DataHandler) BackupDB() error {
    //since this is sqlite, we can simply copy the db file to backup
    return errors.New("Backing up the database is not currently supported")
}

func (s Sqlite3DataHandler) CloseDB() error {
    //since this is sqlite, we can simply copy the db file to backup
    err := (s.db).Close()
    if err != nil {
        return err
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

func (s Sqlite3DataHandler) AddPlayer(player_name string, email string, tagline string, team int) (error) {
    //start a transaction (tx)
    tx, err := (s.db).Begin()
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

func (s Sqlite3DataHandler) GetAllPlayers() ([]Player, error) {
    rows, err := (s.db).Query("select id, name, email, email_md5, tagline, team from players")
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

//COMBINE THIS WITH THE OTHER ONE
func (s Sqlite3DataHandler) GetAllPlayers_display() ([]PlayerDisplay, error) {
    players, err := s.GetAllPlayers()
    if err != nil {
        return nil, err
    }
    players_display := make([]PlayerDisplay, 0)
    for _, player := range players {
        team, err := s.GetTeamByID(player.Team)
        if err != nil {
            return nil, err
        }
        pd := PlayerDisplay{player.Id, player.Name, player.Email, player.Email_md5, player.Tagline, team.Name}
        players_display = append(players_display, pd)
    }
    return players_display, nil
}

func (s Sqlite3DataHandler) GetPlayerByID(id int) (Player, error) {
    stmt, err := (s.db).Prepare("select name, email, email_md5, tagline, team from players where id = ?")
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

func (s Sqlite3DataHandler) DeletePlayer(player_id int) (error) {
    //start a transaction (tx)
    tx, err := (s.db).Begin()
    if err != nil {
        return err
    }
    stmt, err := tx.Prepare("delete from players where id = (?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(player_id)
    if err != nil {
        return err 
    }
    tx.Commit()

    return nil
}

func (s Sqlite3DataHandler) AddGame(offenderA int, defenderA int, offenderB int, defenderB int, scoreA int, scoreB int, winner string, dt string) (error) {
    //start a transaction (tx)
    tx, err := (s.db).Begin()
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

func (s Sqlite3DataHandler) GetAllGames() ([]Game, error) {
    rows, err := (s.db).Query("select id, offenderA, defenderA, offenderB, defenderB, scoreA, scoreB, winner, dt from games")
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

//COMBINE THIS WITH THE OTHER ONE
func (s Sqlite3DataHandler) GetAllGames_display() ([]GameDisplay, error) {
    games, err := s.GetAllGames()
    if err != nil {
        return nil, err
    }
    games_display := make([]GameDisplay, 0)
    for _, game := range games {
        offA, err := s.GetPlayerByID(game.OffenderA)
        if err != nil {
            return nil, err
        }
        defA, err := s.GetPlayerByID(game.DefenderA)
        if err != nil {
            return nil, err
        }
        offB, err := s.GetPlayerByID(game.OffenderB)
        if err != nil {
            return nil, err
        }
        defB, err := s.GetPlayerByID(game.DefenderB)
        if err != nil {
            return nil, err
        }
        gd := GameDisplay{game.Id, offA.Name, defA.Name, offB.Name, defB.Name, game.ScoreA, game.ScoreB, game.Winner, game.Timestamp}
        games_display = append(games_display, gd)
    }
    return games_display, nil
}

func (s Sqlite3DataHandler) GetGameByID(id int) (Game, error) {
    stmt, err := (s.db).Prepare("select offenderA, defenderA, offenderB, defenderB, scoreA, scoreB, winner, dt from games where id = ?")
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

func (s Sqlite3DataHandler) DeleteGame(game_id int) (error) {
    //start a transaction (tx)
    tx, err := (s.db).Begin()
    if err != nil {
        return err
    }
    stmt, err := tx.Prepare("delete from games where id = (?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(game_id)
    if err != nil {
        return err 
    }
    tx.Commit()

    return nil
}

func (s Sqlite3DataHandler) GetAllTeams() ([]Team, error) {
    rows, err := (s.db).Query("select id, name from teams")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    teams := make([]Team, 0)
    for rows.Next() {
        team := Team{}
        rows.Scan(&(team.Id), &(team.Name))
        team.Members, _ = s.GetTeamMembers(team)
        teams = append(teams, team)
    }

    return teams, nil
}

func (s Sqlite3DataHandler) GetTeamByID(id int) (Team, error) {
    stmt, err := (s.db).Prepare("select name from teams where id = ?")
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

    players, err := s.GetTeamMembers(team)
    if err != nil {
        return team, err 
    }

    team.Members = players

    return team, nil
}

func (s Sqlite3DataHandler) GetTeamMembers(team Team) ([]Player, error) {
    stmt, err := (s.db).Prepare("select id, name, email, email_md5, tagline, team from players where team = ?")
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

func (s Sqlite3DataHandler) AddTeam(team_name string) (error) {
    //start a transaction (tx)
    tx, err := (s.db).Begin()
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

func (s Sqlite3DataHandler) DeleteTeam(team_id int) (error) {
    //start a transaction (tx)
    tx, err := (s.db).Begin()
    if err != nil {
        return err
    }
    stmt, err := tx.Prepare("delete from teams where id = (?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(team_id)
    if err != nil {
        return err 
    }
    tx.Commit()

    return nil
}
