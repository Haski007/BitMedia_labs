package users

import (
	// "encoding/json"
	// "io/ioutil"
	"net/http"
)

var currentUserID string

type users struct {
	Users []user `json:"objects"`
}

type user struct {
	ID			string `bson:"id"`
	Email     	string `json:"email" bson:"email"`
	Lname     	string `json:"last_name" bson:"last_name"`
	Country   	string `json:"country" bson:"country"`
	City      	string `json:"city" bson:"city"`
	Gender    	string `json:"gender" bson:"gender"`
	BirthDate 	string `json:"birth_date" bson:"birth_date"`
}

// UserHandler - handle POST and GET queries.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addUser(w, r)
	}
}
