package routers

import (
	"github.com/thiagossjc/redsocial/db"
	"net/http"
)

func DelTwitter(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID ", http.StatusBadRequest)
		return
	}

	err := db.DeleteTwitter(ID, IdUserGlobal)
	if err != nil {
		http.Error(w, "Ocorrió un error al intentar borrar el twitter "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
