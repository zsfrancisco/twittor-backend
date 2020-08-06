package routers

import (
	"encoding/json"
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/models"
	"net/http"
)

/* ConsultRelation checks if there's a relation between 2 users */
func ConsultRelation(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Should send the id parameter", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserID = IDUser
	relation.UserRelationID = ID

	var response models.ResponseConsultRelation

	status, err := db.ConsultRelation(relation)
	if err != nil || status == false {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
