package db

import (
	"github.com/thiagossjc/redsocial/models"
	"golang.org/x/crypto/bcrypt"
)

/*Intento Login*/
func TryLogin(email string, password string) (models.User, bool) {

	user, toFind, _ := UserAlredyExist(email)
	//buscar es falso, usuario existe
	if toFind == false {
		return user, false
	}
	//desncriptar password
	passwordBytes := []byte(password) //slice de bytes
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}
	return user, true
}
