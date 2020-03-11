package games

import (
	"fmt"
	"strconv"
	"net/http"
	
	"../errno"
)

// Handler - handle POST and GET queries with path: "/games".
func Handler(w http.ResponseWriter, r *http.Request) {
	pageNumber, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		errno.PrintError(err)
		return
	} else if r.FormValue("page") == "" {
		pageNumber = 0
	} else if pageNumber + 1 > len(gamesCache) / 20 || pageNumber < 0 {
		errno.PrintError(fmt.Errorf("No such page!"))
		return
	}

	drawTable(gamesCache[pageNumber * 20 : pageNumber * 20 + 20], pageNumber * 20)
}

func drawTable(games []game, id int) {
	fmt.Printf("%-10s%-14.10s%-14.10s%-14.10s%-14.14s\n", "id", "points_gained", "win_status", "game_type", "created");
	for i, game := range games {
		fmt.Printf("%-10d%-14.10s%-14.10s%-14.10s%-14.14s\n", i + id, game.PointsGained, game.WinStatus, game.GameType,game.Created);
	}
}