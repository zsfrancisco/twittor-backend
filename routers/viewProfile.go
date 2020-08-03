package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"net/http"
)

/* ViewProfile allows extract the profile's values */
func ViewProfile(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter ID", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)

	if err != nil {
		http.Error(w, "An error occurred while trying to find the record "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
