package routers

import (
	"encoding/json"
	"net/http"
	"redsocial/models"

	"redsocial/db"
)

//Modificar el Perfil del Usuário
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Dados Incorretos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = db.ModifyRecord(u, IdUserGlobal)
	if err != nil {
		http.Error(w, "Ocorrió un error  al intentar modificar el registro "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated) //Todo correcto ha modificado el usuario
}
