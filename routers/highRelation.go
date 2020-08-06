package routers

import (
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/models"
	"net/http"
)

/* HighRelation saves a relation between users */
func HighRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The parameter id is required", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserID = IDUser
	relation.UserRelationID = ID

	status, err := db.InsertRelation(relation)
	if err != nil {
		http.Error(w, "An error occurred while inserting the relation "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "failed to insert relationship", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
