package users

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"log"
	"strconv"
	"net/http"

	"../database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	pageNumber, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		log.Println(err)
		return
	} else if r.FormValue("page") == "" {
		pageNumber = 0
	}

	var users users

	err = database.UsersCollection.Find(bson.M{}).Skip(pageNumber * 20).Limit(20).All(&users.Users)
	if err != nil {
		log.Println(err)
		return
	}

	drawTable(users.Users)
}

func drawTable(users []user) {
	fmt.Printf("%-9s%-36.32s%-14.10s%-24.20s%-18.14s%-10.6s%-40.40s\n", "id", "email", "last_name", "country", "city", "gender", "birth_date");
	for _, user := range users {
		fmt.Printf("%-9s%-36.32s%-14.10s%-24.20s%-18.14s%-10.6s%-40.40s\n", user.ID, user.Email, user.Lname, user.Country, user.City, user.Gender, user.BirthDate);
	}
}