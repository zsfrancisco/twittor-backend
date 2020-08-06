package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/models"
	"net/http"
)

/* Register is the function to create the user registration in the DB */
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error in the received data "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "You must specify a password of at least 6 characters", 400)
		return
	}

	_, found, _ := db.CheckUserAlreadyExists(user.Email)
	if found == true {
		http.Error(w, "There is already a registered user with that email", 400)
		return
	}

	_, status, err := db.InsertedRecord(user)
	if err != nil {
		http.Error(w, "An error occurred while trying to register the user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "insert user record failed", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
