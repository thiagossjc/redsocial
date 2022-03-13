package db

import "golang.org/x/crypto/bcrypt"

//Function para ecriptar password
func EncriptPassowrd(pass string) (string, error) {
	costo := 8                                                     //significa as passadas, quanto más demora mejor, se puede incripar hasta 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo) //solo acepta y devolve slices de byte
	if err != nil {
		panic(err.Error())
	}
	return string(bytes), err
}
