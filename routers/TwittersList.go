package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"redsocial/db"
)

//LeoTweets
func TwittersList(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe Enviar el Parametro id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe Enviar el Parametro page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page")) //converter alphabetic a intero
	if err != nil {
		http.Error(w, "Debe Enviar el Parametro page con un valor mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(page) //converter a pagina de int a un int64
	//la rotina de bson obriga que isso seja int64

	response, correct := db.SelectTwiiters(ID, pag)
	if !correct {
		http.Error(w, "Error ao leer los twitters", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
