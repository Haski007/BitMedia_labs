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

///////////////////////////// Patterns \\\\\\\\\\\\\\\\\\\\\\\\\\\

// EmailPattern - pattern string to validate user's email!
var EmailPattern = "^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$"

// DatePattern - pattern string to validate user's Birth Date!
var DatePattern = `^\w*,\s\w*\s\d\d?,\s\d{4}\s\d\d?:\d{2}\s[PA]M$`