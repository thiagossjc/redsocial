package handlers

import (
	"log"
	"net/http"
	"os"
	"redsocial/middlew"
	"redsocial/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors" //cors son los permisos que doy a mi api sea acesible de cualquer lugar
)

/*Manejadores seteo mi puerto, el handre y pobo*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/viewpRofile", middlew.CheckDB(middlew.ValidateJWT(routers.Login))).Methods("GET")
	router.HandleFunc("/modifyUser", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/insertTwttet", middlew.CheckDB(middlew.ValidateJWT(routers.InsertTweet))).Methods("POST")
	router.HandleFunc("/listTwitters", middlew.CheckDB(middlew.ValidateJWT(routers.TwittersList))).Methods("GET")
	router.HandleFunc("/removeTwitter", middlew.CheckDB(middlew.ValidateJWT(routers.DelTwitter))).Methods("DELETE")
	router.HandleFunc("/uploadAvatars", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadBanners", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/listUsers", middlew.CheckDB(middlew.ValidateJWT(routers.UsersListRelationship))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/getBanner", middlew.CheckDB(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/addRelationship", middlew.CheckDB(middlew.ValidateJWT(routers.AddRelationship))).Methods("POST")
	router.HandleFunc("/removeRelationship", middlew.CheckDB(middlew.ValidateJWT(routers.RemoveRelationship))).Methods("DELETE")
	router.HandleFunc("/findRelationship", middlew.CheckDB(middlew.ValidateJWT(routers.FindRelationship))).Methods("GET")
	router.HandleFunc("/viewTwittersFollowers", middlew.CheckDB(middlew.ValidateJWT(routers.ViewTwittersFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //Servidor ecucher para ver todos los llamados de petición
}
