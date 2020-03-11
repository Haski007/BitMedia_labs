package games

import (
	"math/rand"
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"

	"../database"
	"../errno"
)

// InitGamesCache - save all game to a variable "gamesCache"
func InitGamesCache(fileName string) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		errno.PrintError(err)
		return
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		errno.PrintError(err)
		return
	}

	var games games

	json.Unmarshal(bytes, &games)
	gamesCache = games.Games
}

// InitTestData - create a new collection "games" at DB "BitMedia" and store
// all test data from games.json file.
func InitTestData() {
	count, err := database.UsersCollection.Count()
	if err != nil {
		errno.PrintError(err)
		return
	}

	var randomGamesArray []int
	
	var userGames userGames

	for i := 0; i < count; i++ {
		randomGamesArray = getRandomGamesArray(gamesCache)
		userGames.ID = i
		userGames.GameIDs = randomGamesArray
		database.UserGamesCollection.Insert(userGames)
		fmt.Printf("#%d - collected!\n", i)
	}
	fmt.Println("Games collection has been stored by test data!")
}

func getRandomGamesArray(gamesCache []game) []int {
	var userGamesArray []int
	var lenUserGamesArray int

	lenUserGamesArray = rand.Intn(100)

	userGamesArray = make([]int, lenUserGamesArray)

	for i := 0; i < lenUserGamesArray; i++ {
		userGamesArray[i] = rand.Intn(len(gamesCache) - 1)
	}
	return userGamesArray
}