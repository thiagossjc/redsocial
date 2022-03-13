package routers

import (
	"github.com/thiagossjc/redsocial/db"
	"io"
	"net/http"
	"os"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar los parametros ID ", http.StatusBadRequest)
		return
	}
	profile, err := db.LokkingProfile(ID)

	if err != nil {
		http.Error(w, "Usuário no encontrato ", http.StatusBadRequest)
		return
	}
	OpenFile, err1 := os.Open("uploads/banners" + profile.Avatar)
	if err1 != nil {
		http.Error(w, "Imagén no encontrato ", http.StatusBadRequest)
		return
	}
	_, err2 := io.Copy(w, OpenFile) //envia una copia del archivo de forma binária
	if err2 != nil {
		http.Error(w, "Error al copiar el imagen ", http.StatusBadRequest)
		return
	}
}
