# API Golang

#### Enter `make` at project folder to compile a binary `BitMedia`

#### To run a program enter `./BitMedia`

#### Configuration file
- `config.go` file path: `conf/conf.go`

```golang
package config

// Port on witch server will be listened.
var Port = ":8080"

// HostName on witch server will be listened.
var HostName = "localhost"

// UsersDataFile - json file name with users test data.
var UsersDataFile = "resources/users_go.json"


///////////////////////////// DataBase \\\\\\\\\\\\\\\\\\\\\\\\\\\

// DataBaseHost, hostname on witch mongoDB is hosting.
var DataBaseHost = "192.168.99.100"

// DataBasePort, port on witch mongoDB is hosting.
var DataBasePort = ":27017"
```

#### Database used `mongoDB`

### Functional
- POST query on path `/user` creates a new user.
`curl -d "data.json" -X POST http://localhost:8080/user`
- GET query on path /users with param `page` draws table of users.
`curl "http://localhost:8080/users?page=23"`
