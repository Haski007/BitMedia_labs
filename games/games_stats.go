package games

import (
	"reflect"
	"github.com/globalsign/mgo/bson"
	"fmt"
	"net/http"
	"regexp"

	"../database"
	"../errno"
)

type statPerDay struct {
	Month	string `bson:"month"`
	Day		string `bson:"day"`
	Year	string `bson:"year"`
	NumOfGames int64 `bson:"num_of_games"`
}

func (day statPerDay) ToString() string {
	result := fmt.Sprintf("%2s\\%02s\\%2s - Count of games: %d\n", 
	day.Month, day.Day, day.Year, day.NumOfGames)
	return result
}

type m bson.M

// StatsHandler - handle GET request to print grop of games!
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("by") {
	case "day":
		getStatPerDay()
	case "gametype":
		getStatPerGameType()
	}
}

// getStatPerDay - print statistic of games per day!
func getStatPerDay() {
	pipeline := []m{
		m{
			"$group" : m{
				"_id" : m{
					"month" : m{"$month" : "$created"},
					"day" : m{"$dayOfMonth" : "$created"},
					"year" : m{"$year" : "$created"},
				},
				"num_of_games" : m{"$sum" : 1}}}}
	pipe := database.GamesCollection.Pipe(pipeline)
	result := []m{}
	err := pipe.All(&result)
	if err != nil {
		errno.PrintError(err)
		return
	}
	var stats []statPerDay

	for i := 0; i < len(result); i++ {
		var sumQ int64
		switch reflect.TypeOf(result[i]["num_of_games"]).Name() {
		case "int":
			sumQ = int64(result[i]["num_of_games"].(int))
		case "int64":
			sumQ = result[i]["num_of_games"].(int64)
		}
		dateString := fmt.Sprintf("%s", result[i]["_id"])
		
		stat := statPerDay{
			Month: regexp.MustCompile(`month:%!s\(int=(\d*)`).FindStringSubmatch(dateString)[1],
			Day: regexp.MustCompile(`day:%!s\(int=(\d*)`).FindStringSubmatch(dateString)[1],
			Year: regexp.MustCompile(`year:%!s\(int=(\d*)`).FindStringSubmatch(dateString)[1],
			NumOfGames: sumQ,
		}
		stats = append(stats, stat)
	}

	for _, stat := range stats {
		fmt.Print(stat.ToString())
	}
}

type statPerGameType struct {
	GameType	string
	NumOfGames 	int64 `bson:"num_of_games"`
}

func (gType statPerGameType) ToString() string {
	result := fmt.Sprintf("Game type (%s) - Count of games: %d\n",
	gType.GameType, gType.NumOfGames)
	return result
}

func getStatPerGameType() {
	pipeline := []m{
		m{
			"$group" : m{
				"_id" : "$game_type",
				"num_of_games" : m{"$sum" : 1}}}}
	pipe := database.GamesCollection.Pipe(pipeline)
	result := []m{}
	err := pipe.All(&result)
	if err != nil {
		errno.PrintError(err)
		return
	}
	var stats []statPerGameType

	for i := 0; i < len(result); i++ {
		var sumQ int64
		switch reflect.TypeOf(result[i]["num_of_games"]).Name() {
		case "int":
			sumQ = int64(result[i]["num_of_games"].(int))
		case "int64":
			sumQ = result[i]["num_of_games"].(int64)
		}
		
		stat := statPerGameType{
			GameType: fmt.Sprintf("%s", result[i]["_id"]),
			NumOfGames: sumQ,
		}
		stats = append(stats, stat)
	}

	for _, stat := range stats {
		fmt.Print(stat.ToString())
	}
}