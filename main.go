package main

import (
	"fmt"
	"net/http"

	"./database"
	"./config"
	"./users"
	"./logger"
)


func main() {
	
	http.HandleFunc("/user", logger.PreLogs(users.UserHandler))
	http.HandleFunc("/users", logger.PreLogs(users.Handler))

	database.InitDB()

	// users.InitTestData(config.UsersDataFile)

	fmt.Println("Listerning on port :8080")

	http.ListenAndServe(config.Port, nil)

}