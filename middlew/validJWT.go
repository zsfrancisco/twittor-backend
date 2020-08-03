package middlew

import (
	"github.com/zsfrancisco/twittor-backend/routers"
	"net/http"
)

/* ValidJWT allows to validate the JWT in the petition */
func ValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _, _, err := routers.ProcessToken(request.Header.Get("Authorization"))
		if err != nil {
			http.Error(writer, "Error in token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
