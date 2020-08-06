package routers

import (
	"github.com/zsfrancisco/twittor-backend/db"
	"io"
	"net/http"
	"os"
)

/* GetBanner sends the banner to http*/
func GetBanner(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Should send the parameter id", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "An error occurred while copying the image", http.StatusBadRequest)
	}
}