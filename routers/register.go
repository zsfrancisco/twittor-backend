package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/models"
	"net/http"
)

/* Register is the function to create the user registration in the DB */
func Register(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Error in the received data "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(writer, "Email is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(writer, "You must specify a password of at least 6 characters", 400)
		return
	}

	_, found, _ := db.CheckUserAlreadyExists(user.Email)
	if found == true {
		http.Error(writer, "There is already a registered user with that email", 400)
		return
	}

	_, status, err := db.InsertedRecord(user)
	if err != nil {
		http.Error(writer, "An error occurred while trying to register the user"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(writer, "insert user record failed", 400)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
