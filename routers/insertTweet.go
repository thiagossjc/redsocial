package routers

import (
	"encoding/json"
	"net/http"
	"redsocial/db"
	"redsocial/models"
	"time"
)

//grabar el twitter en la base de dados
func InsertTweet(w http.ResponseWriter, r *http.Request) {
	var t models.TwitterRequest
	err := json.NewDecoder(r.Body).Decode(&t)

	record := models.Tweet{
		UserID:  IdUserGlobal, //variable global
		Message: t.Message,
		Date:    time.Now(),
	}
	var status bool
	_, status, err = db.InsertTweet(record)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar un Twttet "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
