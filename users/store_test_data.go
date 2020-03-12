package users

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"

	"../database"
	"../errno"
)

// InitUsersCollection - create a new collection "users" at DB "BitMedia" and store
// all test data from users_go.json file.
func InitUsersCollection(fileName string) {
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

	var users users

	json.Unmarshal(bytes, &users)
		
	for i := 0; i < len(users.Users); i++ {
		fmt.Printf("#%d - collected!\n", i)
		users.Users[i].ID = i
		database.UsersCollection.Insert(users.Users[i])
	}
	fmt.Println("Users collection has been stored by test data!")
}