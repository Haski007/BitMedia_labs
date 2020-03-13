package main

import (
	"fmt"
	"net/http"
	
	"./users"
	"./games"
	"github.com/Haski007/BitMedia_labs/config"
	"github.com/Haski007/BitMedia_labs/logger"
	"github.com/Haski007/BitMedia_labs/database"
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