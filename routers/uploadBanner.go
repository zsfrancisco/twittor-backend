package routers

import (
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/models"
	"io"
	"net/http"
	"os"
	"strings"
)

/* UploadBanner uploads the banner to server */
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archive string = "uploads/banners/" + IDUser + "." + extension

	/* opening file, write only and create parameters, 0666 are the permissions,, that returns a file pointer and err */
	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "An error occurred while uploading the image "+err.Error(), http.StatusBadRequest)
		return
	}

	/* Saving in disk */
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "An error occurred while copying the image "+err.Error(), http.StatusBadRequest)
		return
	}

	/* Modifying the user avatar in database */
	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = db.ModifyRegister(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "An error occurred while saving the image in database "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
