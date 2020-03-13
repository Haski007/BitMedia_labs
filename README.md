# API Golang
- Creates three collection at mongoDB and fill them by test data:
  * UserCollection
  * GamesCollection
  * UserGamesCollection
- There is ability:
  * to create a new user;
  * to get list of all users with navigation by pages;
  * to get list of all games with navigation by pages;
  * to get users raiting list with users and their score and rank with navigation by pages;
  * to get game grouped by day/game_type.
  
#### Enter `make` at project folder to compile a binary `BitMedia`

#### To run a program enter `./BitMedia`

#### Configuration file
- `config.go` file path: `conf/conf.go`

```golang
package config

// MaxUserGames max numbers games for one user.
var MaxUserGames = 100

// Port on which server will be listened.
var Port = ":8080"

// HostName on which server will be listened.
var HostName = "localhost"

// UsersDataFile - json file name with users test data.
var UsersDataFile = "resources/users_go.json"

// GamesDataFile - json file name with games test data.
var GamesDataFile = "resources/games.json"


///////////////////////////// DataBase \\\\\\\\\\\\\\\\\\\\\\\\\\\

// DataBaseHost - hostname on which mongoDB is hosting.
var DataBaseHost = "192.168.99.101"

// DataBasePort - port on which mongoDB is hosting.
var DataBasePort = ":27017"
```

#### Database used `mongoDB`

### Functional
- POST query on path `/user` creates a new user.
```bash
curl -d "data.json" -X POST http://localhost:8080/user
```
- GET query on path `/users` with param `page` draws table of all users.
```bash
curl "http://localhost:8080/users?page=21"
```

- GET query on path `/games` with param `page` draws table of all games.
```bash
curl "http://localhost:8080/games?page=42"
```

- GET query on path `/games/stats` with param `by` which can be (`day/game_type`).
```bash
curl "http://localhost:8080/games/stats?by=game_type"
```

- GET query on path `/users/rating` with param `page` draws rating table of users.
```bash
curl "http://localhost:8080/users/rating?page=81"
```

