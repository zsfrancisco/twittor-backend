package db

import (
	"github.com/zsfrancisco/twittor-backend/models"
	"golang.org/x/crypto/bcrypt"
)

/* TryLogin checks the user and gives access to db */
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserAlreadyExists(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
