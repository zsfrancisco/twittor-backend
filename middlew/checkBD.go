package middlew

import (
	"github.com/zsfrancisco/twittor-backend/db"
	"net/http"
)

/* CheckBD is the middleware to know the database status */
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(writer, "The connection to the database was lost", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
