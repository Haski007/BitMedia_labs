package main

import (
	"fmt"
	"net/http"

	"./database"
	"./config"
	"./users"
	"./logger"
	"./games"
)

func main() {
	
	http.HandleFunc("/user", logger.PreLogs(users.UserHandler))
	http.HandleFunc("/users", logger.PreLogs(users.Handler))
	http.HandleFunc("/games", logger.PreLogs(games.Handler))

	database.InitDB()

	// users.InitTestData(config.UsersDataFile)
	games.InitGamesCache(config.GamesDataFile)
	// games.InitTestData()

	fmt.Println("Listerning on port :8080")

	http.ListenAndServe(config.Port, nil)

}