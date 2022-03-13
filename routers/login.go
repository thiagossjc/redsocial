package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/thiagossjc/redsocial/db"
	jwt "github.com/thiagossjc/redsocial/jwt"
	"github.com/thiagossjc/redsocial/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Usuário y/o contraseñas inválidos! "+err.Error(), 400)
		return //Cancela Endpoint
	}

	if len(u.Email) == 0 {
		http.Error(w, "El email de usuário é requerido!", 400) //aqui no hay error
		return                                                 //Cancela Endpoint
	}

	document, exist := db.TryLogin(u.Email, u.Password)
	if !exist {
		http.Error(w, "Isuário e contraseña inválido!", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document) //Irá receber el document y devolve untoken
	if err != nil {
		http.Error(w, "Error al intentar generar el Token! "+err.Error(), 400)
		return //Cancela Endpoint
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //Devolver status ok
	json.NewEncoder(w).Encode(resp)

	//Como grabar la cookie del usuario
	expirationTimeCookie := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTimeCookie,
	})
}
