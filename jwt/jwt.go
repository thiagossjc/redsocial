package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/thiagossjc/redsocial/models"
)

//Generar Token
func GenerateJWT(u models.User) (string, error) {
	myKey := []byte("Master del Desarrollo Engrenelog") //Criamos la llave privada
	//Criar la lista de privilégios
	payload := jwt.MapClaims{
		"email":      u.Email,
		"name":       u.Name,
		"last_name":  u.LastName,
		"date_birth": u.DateBirth,
		"biography":  u.Biography,
		"web_site":   u.WebSite,
		"_id":        u.Id.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //Criamos un Token JWT
	tokenStr, err := token.SignedString(myKey)                  //La firma de desarrollo

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
