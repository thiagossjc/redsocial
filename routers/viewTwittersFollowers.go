package routers

import (
	"encoding/json"
	"net/http"
	"redsocial/db"
	"strconv"
)

func ViewTwittersFollowers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro page ", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro page mayor que 1", http.StatusBadRequest)
		return
	}
	response, correct := db.SelectTwittersFollowers(IdUserGlobal, page)

	if !correct {
		http.Error(w, "Error al leer los twitters ", http.StatusBadRequest)
	}
	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
