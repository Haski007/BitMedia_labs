package users

import (
	"net/http"
)

var currentUserID int

type users struct {
	Users []user `json:"objects"`
}

type user struct {
	ID			int    `json:"_id" bson:"_id"`
	Email     	string `json:"email" bson:"email"`
	Lname     	string `json:"last_name" bson:"last_name"`
	Country   	string `json:"country" bson:"country"`
	City      	string `json:"city" bson:"city"`
	Gender    	string `json:"gender" bson:"gender"`
	BirthDate 	string `json:"birth_date" bson:"birth_date"`
}

// UserHandler - handle POST and GET queries with path: "/user".
func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addUser(w, r)
	}
}
