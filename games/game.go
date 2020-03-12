package games

import (
	"time"
)

var gamesCache []game


type userGames struct {
	ID      int   `bson:"_id"`
	GameIDs []int `bson:"game_ids"`
}

type games struct {
	Games []game `json:"objects"`
}

type game struct {
	ID           int    `bson:"_id"`
	PointsGained string `json:"points_gained" bson:"points_gained"`
	WinStatus    string `json:"win_status" bson:"win_status"`
	GameType     string `json:"game_type" bson:"game_type"`
	CreatedStr   string `json:"created" bson:"-"`
	CreatedISODate time.Time `bson:"created"`
}