package middlew

import (
	"net/http"

	"github.com/thiagossjc/redsocial/db"
)

//middlew siempre tinen que devolver el mismo tipo de dado que recebió
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { //retorna funcion anonima
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexión Perdida con la Base de Dados!", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
