package main

import (
  "html/template"
  "github.com/gorilla/schema"
)

// html form decoder
var decoder = schema.NewDecoder()

var index_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/index.html",
))
var players_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/players.html",
))
var addplayer_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/add_player.html",
))
var deleteplayer_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/delete_player.html",
))
var player_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/player.html",
))
var games_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/games.html",
))
var addgame_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/add_game.html",
))
var deletegame_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/delete_game.html",
))
var game_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/game.html",
))
var teams_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/teams.html",
))
var addteam_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/add_team.html",
))
var deleteteam_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/delete_team.html",
))
var team_html = template.Must(template.ParseFiles(
    "templates/_base.html",
    "templates/team.html",
))
var queue_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/queue.html",
))
var rankings_html = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/rankings.html",
))


