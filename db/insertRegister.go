package db

import (
	"context"
	"time"

	"github.com/thiagossjc/redsocial/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Insert Registro de User
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //agregar contexto miniatura para el tem de timeout
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")
	u.Password, _ = EncriptPassowrd(u.Password)
	result, err := col.InsertOne(ctx, u) //Inserindo en la colection
	if err != nil {
		return "", false, err //el primer es un id vazio caso de error
	}
	objId, _ := result.InsertedID.(primitive.ObjectID) //Buscando elID da inserção en el object
	return objId.String(), true, nil                   //Convertendo el objecto en una string
}
