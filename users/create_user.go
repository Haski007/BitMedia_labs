package users

import (
	"strconv"
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"

	"../database"
)

func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser user

	currentUserID = getCurrentUserID()
	
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	
	err = json.Unmarshal(bytes, &newUser)
	if err != nil {
		log.Println(err)
		return
	}

	err = validateUser(newUser)
	if err != nil {
		log.Println(err)
		return
	}

	newUser.ID = currentUserID

	err = database.UsersCollection.Insert(newUser)
	if err != nil {
		log.Println(err)
		return
	}

	successCreatingPrint(newUser)
}

func getCurrentUserID() string {
	count, err := database.UsersCollection.Count()
	if err != nil {
		return "0"
	}
	return strconv.Itoa(count)
}

func successCreatingPrint(newUser user) {
	fmt.Printf("Successfully created user #%s:\n", newUser.ID)
	fmt.Printf("Email: %s\n", newUser.Email)
	fmt.Printf("Last Name: %s\n", newUser.Lname)
	fmt.Printf("Country: %s\n", newUser.Country)
	fmt.Printf("City: %s\n", newUser.City)
	fmt.Printf("Gender: %s\n", newUser.Gender)
	fmt.Printf("Birth Date: %s\n", newUser.BirthDate)
}