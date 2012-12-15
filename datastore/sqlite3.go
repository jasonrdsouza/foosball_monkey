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
    "path/filepath"
    "io/ioutil"
)

type Sqlite3DataHandler struct {
    db_name string
    db *sql.DB
}

func (s *Sqlite3DataHandler) CreateNewDB(db_name string) error {
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

func (s *Sqlite3DataHandler) ConnectToDB(db_name string) error {
    s.db_name = "./" + db_name

    temp_db, err := sql.Open("sqlite3", s.db_name)
    s.db = temp_db
    if err != nil {
        return err
    }
    return nil
}

func (s *Sqlite3DataHandler) BackupDB(backup_dir string) error {
    //since this is sqlite, we can simply copy the db file to backup
    db_contents, err := ioutil.ReadFile(s.db_name)
    if err != nil {
        return err
    }
    backup_filepath := filepath.Join(backup_dir, s.db_name)
    err = ioutil.WriteFile(backup_filepath, db_contents, 0644)
    if err != nil {
        return err
    }
    return nil
}

func (s Sqlite3DataHandler) CloseDB() error {
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

func (s *Sqlite3DataHandler) AddPlayer(player_name string, email string, tagline string, team int) (error) {
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

func (s *Sqlite3DataHandler) GetAllPlayers() ([]Player, error) {
    rows, err := (s.db).Query("SELECT p.id, p.name, p.email, p.email_md5, p.tagline, p.team, t.name FROM players p JOIN teams t ON p.team = t.id")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    players := make([]Player, 0)
    for rows.Next() {
        player := Player{}
        rows.Scan(&(player.Id), &(player.Name), &(player.Email), &(player.Email_md5), &(player.Tagline), &(player.Team_id), &(player.Team))
        players = append(players, player)
    }

    return players, nil
}

func (s *Sqlite3DataHandler) GetPlayerByID(id int) (Player, error) {
    stmt, err := (s.db).Prepare("SELECT p.name, p.email, p.email_md5, p.tagline, p.team, t.name FROM players p JOIN teams t ON p.team = t.id WHERE p.id = ?")
    if err != nil {
        return Player{}, err
    }
    defer stmt.Close()

    player := Player{}
    player.Id = id
    err = stmt.QueryRow(id).Scan(&(player.Name), &(player.Email), &(player.Email_md5), &(player.Tagline), &(player.Team_id), &(player.Team))
    if err != nil {
        return Player{}, err 
    }

    return player, nil
}

func (s *Sqlite3DataHandler) DeletePlayer(player_id int) (error) {
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

func (s *Sqlite3DataHandler) AddGame(offenderA int, defenderA int, offenderB int, defenderB int, scoreA int, scoreB int, winner string, dt string) (error) {
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

func (s *Sqlite3DataHandler) GetAllGames() ([]Game, error) {
    rows, err := (s.db).Query("SELECT g.id, g.offenderA, g.defenderA, g.offenderB, g.defenderB, g.scoreA, g.scoreB, g.winner, g.dt FROM games g")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    games := make([]Game, 0)
    for rows.Next() {
        game := Game{}
        var date_string string
        rows.Scan(&(game.Id), &(game.OffenderA_id), &(game.DefenderA_id), &(game.OffenderB_id), &(game.DefenderB_id), &(game.ScoreA), &(game.ScoreB), &(game.Winner), &(date_string))
        game.Timestamp, _ = convertDateStrToTime(date_string)
        games = append(games, game)
    }

    games2 := make([]Game, 0)
    for _, game := range games {
        offA, err := s.GetPlayerByID(game.OffenderA_id)
        if err != nil {
            return nil, err
        }
        defA, err := s.GetPlayerByID(game.DefenderA_id)
        if err != nil {
            return nil, err
        }
        offB, err := s.GetPlayerByID(game.OffenderB_id)
        if err != nil {
            return nil, err
        }
        defB, err := s.GetPlayerByID(game.DefenderB_id)
        if err != nil {
            return nil, err
        }
        g := Game{game.Id, game.OffenderA_id, game.DefenderA_id, game.OffenderB_id, game.DefenderB_id, offA.Name, defA.Name, offB.Name, defB.Name, game.ScoreA, game.ScoreB, game.Winner, game.Timestamp}
        games2 = append(games2, g)
    }

    return games2, nil
}

func (s *Sqlite3DataHandler) GetGameByID(id int) (Game, error) {
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

func (s *Sqlite3DataHandler) DeleteGame(game_id int) (error) {
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

func (s *Sqlite3DataHandler) GetAllTeams() ([]Team, error) {
    rows, err := (s.db).Query("select id, name from teams")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    teams := make([]Team, 0)
    for rows.Next() {
        team := Team{}
        rows.Scan(&(team.Id), &(team.Name))
        team.Members, err = s.GetTeamMembers(team)
        if err != nil {
            return nil, err
        }
        teams = append(teams, team)
    }

    return teams, nil
}

func (s *Sqlite3DataHandler) GetTeamByID(id int) (Team, error) {
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

func (s *Sqlite3DataHandler) GetTeamMembers(team Team) ([]Player, error) {
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

func (s *Sqlite3DataHandler) AddTeam(team_name string) (error) {
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

func (s *Sqlite3DataHandler) DeleteTeam(team_id int) (error) {
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
