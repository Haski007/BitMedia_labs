package users

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"

	"../database"
	"github.com/Haski007/BitMedia_labs/errno"
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

	var usersI []interface{}
	for i, t := range users.Users {
		t.ID = i
		usersI = append(usersI, t)
	}

	database.UsersCollection.Insert(usersI...)

	fmt.Println("Users collection has been stored by test data!")
}