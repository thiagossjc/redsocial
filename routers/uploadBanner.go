package routers

import (
	"io"
	"net/http"
	"os"
	"redsocial/db"
	"redsocial/models"
	"strings"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, _ := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archive string = "uploads/banners" + IdUserGlobal + "." + extension
	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir el banner "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file) //Gravo el arquivo en disco

	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	var status bool
	user.Banner = IdUserGlobal + "-" + extension
	status, err1 := db.ModifyRecord(user, IdUserGlobal)

	if err1 != nil || status == false {
		http.Error(w, "Error al grabar el Banner en la BD "+err1.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
