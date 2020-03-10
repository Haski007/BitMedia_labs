package logger

import (
	"log"
	"net/http"
)

// PreLogs writes to logs method of response and ip of user and set json header.
func PreLogs(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server [net/http] method %s connection from [%v]", r.Method, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}