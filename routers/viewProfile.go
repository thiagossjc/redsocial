package routers

import (
	"encoding/json"
	"net/http"

	"github.com/thiagossjc/redsocial/db"
)

//Ver Perfil permite extraer los valores del Perfil
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	profile, err := db.LokkingProfile(ID)
	if err != nil {
		http.Error(w, "Ocorrió un error a intentar buscar el retristro "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "applicaation/json")
	w.WriteHeader(http.StatusCreated) //Devolverá status 202, satisfatório
	json.NewEncoder(w).Encode(profile)

}
