package games

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/Haski007/BitMedia_labs/database"
	"github.com/Haski007/BitMedia_labs/errno"
	"github.com/Haski007/BitMedia_labs/config"
)

// InitGamesCollection - save all game to a variable "gamesCache"
func InitGamesCollection(fileName string) {
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

	var gamesI []interface{}
	for i, game := range gamesCache {
		game.ID = i
		game.CreatedISODate, err = time.Parse("1/2/2006 3:04 PM", game.CreatedStr)
		if err != nil {
			errno.PrintError(err)
			return
		}
		gamesI = append(gamesI, game)
	}

	database.GamesCollection.Insert(gamesI...)

	fmt.Println("Games collection has been stored by test data!")
	
}

// InitUserGamesCollection - create a new collection "games" at DB "BitMedia" and store
// all test data from games.json file.
func InitUserGamesCollection() {
	count, err := database.UsersCollection.Count()
	if err != nil {
		errno.PrintError(err)
		return
	}

	var randomGamesArray []int
	
	var userGames userGames

	var userGamesI []interface{}

	for i := 0; i < count; i++ {
		randomGamesArray = getRandomGamesArray(gamesCache)
		userGames.ID = i
		userGames.GameIDs = randomGamesArray

		userGamesI = append(userGamesI, userGames)
	}
	
	database.UserGamesCollection.Insert(userGamesI...)
	fmt.Println("UserGamesCollection has been stored by test data!")
}

func getRandomGamesArray(gamesCache []game) []int {
	var userGamesArray []int
	var lenUserGamesArray int

	lenUserGamesArray = rand.Intn(config.MaxUserGames)

	userGamesArray = make([]int, lenUserGamesArray)

	for i := 0; i < lenUserGamesArray; i++ {
		userGamesArray[i] = rand.Intn(len(gamesCache) - 1)
	}
	return userGamesArray
}