package database

import (
	"log"

	"../config"
	"github.com/globalsign/mgo"
)

// UsersCollection - users table.
var UsersCollection *mgo.Collection

// GamesCollection - games table.
var UserGamesCollection *mgo.Collection

// InitDB, initialises mongoDB with your configurations at config/config.go.
func InitDB() {
	session, err := mgo.Dial(config.DataBaseHost)
	if err != nil {
		log.Fatal(err)
	}

	UsersCollection = session.DB("BitMedia").C("users")
	UserGamesCollection = session.DB("BitMedia").C("user_games")


	if err = session.Ping(); err != nil {
		log.Fatal(err)
	}
}