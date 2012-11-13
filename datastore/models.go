package datastore

import (
    "time"
)

type Player struct {
    Id int
    Name string
    Email string
    Email_md5 string
    Tagline string
    Team int
}

type PlayerDisplay struct {
    Id int
    Name string
    Email string
    Email_md5 string
    Tagline string
    Team string
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