package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/thiagossjc/redsocial/db"
	"github.com/thiagossjc/redsocial/models"
)

//Email valor de Email usa en todos los endpoints
var Email string

//IdUser valor de IdUser usa en todos los endpoints
var IdUserGlobal string

//Processa token para extrair sus valores
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myPassword := []byte("Master del Desarrollo Engrenelog")
	claims := &models.Claim{}                 //tiene que ser puntero
	splitToken := strings.Split(tk, "Bearer") //converterá o token en un vetor, en el elemento zero tenemos la palabra bearer
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPassword, nil
	})
	if err == nil {
		_, Find, _ := db.UserAlredyExist(claims.Email)
		if Find {
			Email = claims.Email
			IdUserGlobal = string(claims.ID.Hex())
		}
		return claims, Find, IdUserGlobal, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}
	return claims, false, string(""), err
}
