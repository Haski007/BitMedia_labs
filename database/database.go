package database

import (
	"log"

	"../config"
	"github.com/globalsign/mgo"
)

// UsersCollection - users table.
var UsersCollection *mgo.Collection

// GamesCollection - games table.
var GamesCollection *mgo.Collection

// InitDB, initialises mongoDB with your configurations at config/config.go.
func InitDB() {
	session, err := mgo.Dial(config.DataBaseHost)
	if err != nil {
		log.Fatal(err)
	}

	UsersCollection = session.DB("BitMedia").C("users")
	GamesCollection = session.DB("BitMedia").C("games")


	if err = session.Ping(); err != nil {
		log.Fatal(err)
	}
}