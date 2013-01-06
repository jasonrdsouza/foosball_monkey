package datastore

import (
    //"appengine"
    //"appengine/datastore"
)

type AppengineDataHandler struct {
    db_name string
}

func (a AppengineDataHandler) CreateNewDB(db_name string) error {
    return nil
}

func (a AppengineDataHandler) ConnectToDB(db_name string) error {
    return nil
}

func (a AppengineDataHandler) BackupDB(backup_dir string) error {
    return nil
}

func (a AppengineDataHandler) CloseDB() error {
    return nil
}

func (a AppengineDataHandler) AddPlayer(player_name, email, tagline string, team int) error {
    return nil
}

func (a AppengineDataHandler) DeletePlayer(player_id int) error {
    return nil
}

func (a AppengineDataHandler) GetAllPlayers() ([]Player, error) {
    return nil, nil
}

func (a AppengineDataHandler) GetPlayerByID(id int) (Player, error) {
    return Player{}, nil
}

func (a AppengineDataHandler) AddGame(offenderA, defenderA, offenderB, defenderB, scoreA, scoreB int, winner, dt string) error {
    return nil
}

func (a AppengineDataHandler) DeleteGame(game_id int) error {
    return nil
}

func (a AppengineDataHandler) GetAllGames() ([]Game, error) {
    return nil, nil
}

func (a AppengineDataHandler) GetGameByID(id int) (Game, error) {
    return Game{}, nil
}

func (a AppengineDataHandler) AddTeam(team_name string) error {
    return nil
}

func (a AppengineDataHandler) DeleteTeam(team_id int) error {
    return nil
}

func (a AppengineDataHandler) GetAllTeams() ([]Team, error) {
    return nil, nil
}

func (a AppengineDataHandler) GetTeamByID(id int) (Team, error) {
    return Team{}, nil
}

func (a AppengineDataHandler) GetTeamMembers(team Team) ([]Player, error) {
    return nil, nil
}
