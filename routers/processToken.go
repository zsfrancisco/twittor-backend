package routers

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/models"
	"strings"
)

/* Email is the email value used in all of the endpoints */
var Email string

/* IDUser is the id returned from the model, it's used in all of the endpoints */
var IDUser string

/* ProcessToken process the token to extract its values */
func ProcessToken(token string) (*models.Claim, bool, string, error)  {
	myKey := []byte("ClaveDeSeguridadParaCursoTwitter")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := db.CheckUserAlreadyExists(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
