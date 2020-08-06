package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"net/http"
	"strconv"
)

/* ReadRelationTweets reads the tweets of all of our followers */
func ReadRelationTweets(w http.ResponseWriter, r *http.Request)  {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Should send the parameter page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Should send the parameter page as integer greater than 0", http.StatusBadRequest)
		return
	}

	response, correct := db.ReadFollowersTweets(IDUser, page)
	if correct == false {
		http.Error(w, "An error occurred while reading the tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
