package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"net/http"
	"strconv"
)

/* ReadTweets reads the tweets */
func ReadTweets(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Should send the parameter id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Should send the parameter page", http.StatusBadRequest)
		return
	}

	/* atoi converts a string in number */
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Should send the parameter page greater than zero", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	response, correct := db.ReadTweets(ID, pag)
	if correct == false {
		http.Error(w, "An error occurred while reading tweets, try again", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
