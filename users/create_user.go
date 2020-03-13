package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/Haski007/BitMedia_labs/database"
	"github.com/Haski007/BitMedia_labs/errno"
)

func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser user

	currentUserID = getCurrentUserID()
	
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errno.PrintError(err)
		return
	}
	
	err = json.Unmarshal(bytes, &newUser)
	if err != nil {
		errno.PrintError(err)
		return
	}

	err = validateUser(newUser)
	if err != nil {
		errno.PrintError(err)
		return
	}

	newUser.ID = currentUserID


	err = database.UsersCollection.Insert(newUser)
	if err != nil {
		errno.PrintError(err)
		return
	}

	successCreatingPrint(newUser)
}

func getCurrentUserID() int {
	count, err := database.UsersCollection.Count()
	if err != nil {
		return 0
	}
	var u user
	err = database.UsersCollection.Find(nil).Select(bson.M{"_id": 1}).Skip(count - 1).One(&u)
	if err != nil {
		errno.PrintError(err)
		return 0
	}
	return u.ID + 1
}

func successCreatingPrint(newUser user) {
	fmt.Printf("Successfully created user #%d:\n", newUser.ID)
	fmt.Printf("Email: %s\n", newUser.Email)
	fmt.Printf("Last Name: %s\n", newUser.Lname)
	fmt.Printf("Country: %s\n", newUser.Country)
	fmt.Printf("City: %s\n", newUser.City)
	fmt.Printf("Gender: %s\n", newUser.Gender)
	fmt.Printf("Birth Date: %s\n", newUser.BirthDate)
}