package routers

import (
	"encoding/json"
	"net/http"
	"redsocial/db"
	"strconv"
)

//Lista de Usuários
func UsersListRelationship(w http.ResponseWriter, r *http.Request) {

	TypeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("page")
	pagTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Debe enviar el parámetro page como entero mayor a cero", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)
	results, status := db.SelectAllUsers(IdUserGlobal, pag, search, TypeUser)

	if !status {
		http.Error(w, "Error al leer los usuarios ", http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}
