package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/models"
	"net/http"
	"time"
)

/* SaveTweet allows save the tweet in the database */
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	/* tweet */
	register := models.SaveTweet{
		UserID: IDUser,
		Message: message.Message,
		Date: time.Now(),
	}

	var status bool
	_, status, err = db.InsertTweet(register)
	if err != nil {
		http.Error(w, "An error occurred while inserting the tweet, try again "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "The tweet isn't registered", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
