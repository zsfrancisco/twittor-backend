package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"net/http"
	"strconv"
)

/* ListUsers reads the user list*/
func ListUsers(w http.ResponseWriter, r *http.Request)  {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Should send the parameter page as integer and greater than zero", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := db.ReadAllUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "An error occurred while reading the users",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
