package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/jwt"
	"github.com/zsfrancisco/twittor-backend/models"
	"net/http"
)

/* Login does the login to the system */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "User or password invalid "+err.Error(), 400)
		return
	}

	if len(user.Email)==0 {
		http.Error(w, "email is required", 400)
		return
	}

	document, exist := db.TryLogin(user.Email, user.Password)

	if exist == false {
		http.Error(w, "User and/or password invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "An error occurred while trying to generate the corresponding token "+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
