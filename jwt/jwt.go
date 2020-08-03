package jwt

import (
	"github.com/zsfrancisco/twittor-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

/* GenerateJWT generates a token */
func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("ClaveDeSeguridadParaCursoTwitter")
	payload := jwt.MapClaims{
		"email": user.Email,
		"name": user.Name,
		"surname": user.Surname,
		"birthday": user.Birthday,
		"biography": user.Biography,
		"location": user.Location,
		"website": user.Website,
		"_id": user.ID.Hex(),
		/* Unix formatted as long and is very fast */
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	/* signing token */
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr,nil
}
