package routers

import (
	"encoding/json"
	"net/http"

	"github.com/thiagossjc/redsocial/db"

	"github.com/thiagossjc/redsocial/models"
)

/*Funccion para criar en la base de dados Register de Usuarios*/
func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u) //Body es un Stream = dado solo se puede lerr una vez
	if err != nil {
		http.Error(w, "Error en los dados recibidos "+err.Error(), 400)
		return
	}
	//len te trae el largo de una string
	if len(u.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}
	if len(u.Password) < 6 {
		http.Error(w, "Debe espeficar una contraseña de al menos 6 caracterers", 400)
		return
	}

	_, Found, _ := db.UserAlredyExist(u.Email) //Checará se o usuário já está carregado en la base de dados
	if Found {
		http.Error(w, "Ya existe un usuario registrado con este email!", 400)
		return
	}

	_, status, err := db.InsertRegister(u)
	if err != nil {
		http.Error(w, "Ocurrió un error a intenter realizar el registro de Usuario", 400)
		return
	}

	if !status { //se no ha retornada nada
		http.Error(w, "No se ha logrado insertar el registro de Usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated) //Fue creado con suceso!

}
