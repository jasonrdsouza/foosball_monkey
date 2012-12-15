package datastore

import (
    "time"
)


type FoosballMonkeyDataHandler interface {
    CreateNewDB(db_name string) error
    ConnectToDB(db_name string) error
    BackupDB(backup_dir string) error
    CloseDB() error
    AddPlayer(player_name, email, tagline string, team int) error
    DeletePlayer(player_id int) error
    GetAllPlayers() ([]Player, error)
    GetPlayerByID(id int) (Player, error)
    AddGame(offenderA, defenderA, offenderB, defenderB, scoreA, scoreB int, winner, dt string) error
    DeleteGame(game_id int) error
    GetAllGames() ([]Game, error)
    GetGameByID(id int) (Game, error)
    AddTeam(team_name string) error
    DeleteTeam(team_id int) error
    GetAllTeams() ([]Team, error)
    GetTeamByID(id int) (Team, error)
    GetTeamMembers(team Team) ([]Player, error)
}

type Player struct {
    Id int
    Name string
    Email string
    Email_md5 string
    Tagline string
    Team_id int
    Team string
}

type Game struct {
    Id int
    OffenderA_id int
    DefenderA_id int
    OffenderB_id int
    DefenderB_id int
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