package main

import (
	"fmt"
	"net/http"


	"github.com/Haski007/BitMedia_labs/config"
	"./users"
	"./logger"
	"./games"
	"./database"
)

func main() {
	
	http.HandleFunc("/user", logger.PreLogs(users.UserHandler))
	http.HandleFunc("/users", logger.PreLogs(users.Handler))
	http.HandleFunc("/users/rating", logger.PreLogs(users.RatingHandler))
	http.HandleFunc("/games", logger.PreLogs(games.Handler))
	http.HandleFunc("/games/stats", logger.PreLogs(games.StatsHandler))

	database.InitDB()

	users.InitUsersCollection(config.UsersDataFile)
	games.InitGamesCollection(config.GamesDataFile)
	games.InitUserGamesCollection()

	fmt.Println("Listerning on port :8080")

	http.ListenAndServe(config.Port, nil)

}