package routers

import (
	"encoding/json"
	"net/http"
	"redsocial/db"
	"redsocial/models"
)

func FindRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var relat models.Relationship
	relat.UserId = IdUserGlobal
	relat.UserRelationshipId = ID

	var resp models.FindRelationship

	status, err := db.SelectRelationShip(relat)

	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
