package routers

import (
	"net/http"

	"redsocial/db"
	"redsocial/models"
)

func AddRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro id es obligatório ", http.StatusBadRequest)
		return
	}
	var relation models.Relationship
	relation.UserId = IdUserGlobal
	relation.UserRelationshipId = ID

	status, err := db.InsertRelationship(relation)
	if err != nil || !status {
		http.Error(w, "Erro al intentar crear la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
