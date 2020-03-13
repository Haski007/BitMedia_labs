package users

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"strconv"
	"net/http"

	"github.com/Haski007/BitMedia_labs/database"
	"github.com/Haski007/BitMedia_labs/errno"
)

// Handler - handle POST and GET queries with path: "/users".
func Handler(w http.ResponseWriter, r *http.Request) {
	count, err := database.UsersCollection.Count()
	if err != nil {
		errno.PrintError(err)
		return
	}

	pageNumber, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		errno.PrintError(err)
		return
	} else if r.FormValue("page") == "" {
		pageNumber = 0
	} else if pageNumber + 1 > count / 20 {
		errno.PrintError(fmt.Errorf("No such page!"))
		return
	}

	var users users

	err = database.UsersCollection.Find(bson.M{}).Skip(pageNumber * 20).Limit(20).All(&users.Users)
	if err != nil {
		errno.PrintError(err)
		return
	}

	drawTable(users.Users)
}

func drawTable(users []user) {
	fmt.Printf("%-9s%-36.32s%-14.10s%-24.20s%-18.14s%-10.6s%-40.40s\n", "id", "email", "last_name", "country", "city", "gender", "birth_date");
	for _, user := range users {
		fmt.Printf("%-9d%-36.32s%-14.10s%-24.20s%-18.14s%-10.6s%-40.40s\n", user.ID, user.Email, user.Lname, user.Country, user.City, user.Gender, user.BirthDate);
	}
}