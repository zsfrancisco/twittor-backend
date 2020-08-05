package routers

import (
	"github.com/zsfrancisco/twittor-backend/db"
	"net/http"
)

/* DeleteTweet allows delete a specific tweet */
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Should send the id parameter", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)

	if err != nil {
		http.Error(w, "An error occurred while deleting, try later "+ err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
