package users

import (
	"reflect"
	"github.com/globalsign/mgo/bson"
	"fmt"
	"strconv"
	"net/http"

	"../database"
	"github.com/Haski007/BitMedia_labs/errno"
)

type m bson.M

type userRating struct {
	ID 	int
	Lname string
	BirthDate string
	Rating int64
}


// RatingHandler - handle GET queries with path "/users/rating".
func RatingHandler(w http.ResponseWriter, r *http.Request) {
	count, err := database.UserGamesCollection.Count()
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
		errno.PrintError(fmt.Errorf("No such page"))
		return
	}

	usersRating, err := makeRating(count)
	if err != nil {
		errno.PrintError(err)
		return
	}
	
	fmt.Printf("%-10s%-20.16s%-40.36s%-10s\n", "Rank", "Last Name", "Birth date", "Score")
	for i := pageNumber * 20; i < pageNumber * 20 + 20; i++ {
		fmt.Printf("%-10d%-20.16s%-40.36s%-10d\n", usersRating[i].ID, usersRating[i].Lname, usersRating[i].BirthDate, usersRating[i].Rating)
	}
}

func makeRating(count int) ([]userRating, error) {

	pipeline := []m{
		m{
			"$project" : m{
				"count" : m{
					"$size" : "$game_ids"}}},
		{
			"$sort" : m{
				"count" : -1}}}
	pipe := database.UserGamesCollection.Pipe(pipeline)
	result := []m{}
	err := pipe.All(&result)
	if err != nil {
		return nil, err
	}
	
	var usersRating []userRating

	var users []user

	err = database.UsersCollection.Find(nil).All(&users)
	if err != nil {
		return nil, err
	}

	for i := 0; i < count; i++ {
		var countG int64
		switch reflect.TypeOf(result[i]["count"]).Name() {
		case "int":
			countG = int64(result[i]["count"].(int))
		case "int64":
			countG = result[i]["count"].(int64)
		}
		
		var id int
		switch reflect.TypeOf(result[i]["_id"]).Name() {
		case "int":
			id = int(result[i]["_id"].(int))
		case "int64":
			id = result[i]["_id"].(int)
		}


		userR := userRating{
			ID: i,
			Rating: countG,
			Lname: users[id].Lname,
			BirthDate: users[id].BirthDate,
		}
		usersRating = append(usersRating, userR)
	}

	return usersRating, nil
}