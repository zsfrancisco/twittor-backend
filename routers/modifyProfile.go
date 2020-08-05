package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/models"
	"net/http"
)

/* ModifyProfile modifies the user profile */
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Incorrect data "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.ModifyRegister(user, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while modifying the profile, try again "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "The user registry has not been modified", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
