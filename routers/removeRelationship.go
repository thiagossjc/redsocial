package routers

import (
	"net/http"
	"redsocial/db"

	"github.com/thiagossjc/redsocial/models"
)

func RemoveRelationship(w http.ResponseWriter, r *http.Request) {
	IDR := r.URL.Query().Get("id")
	var relat models.Relationship
	relat.erId = IdUserGlobal
	relat.UserRelationshipId = IDR
	status, err := db.DeleteRelationship(relat)
	if err != nil || status == false {
		http.Error(w, "Informaciones incorrectas para deleción"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
