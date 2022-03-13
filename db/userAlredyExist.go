package db

import (
	"context"
	"time"

	"github.com/thiagossjc/redsocial/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Checar se existe User
func UserAlredyExist(email string) (models.User, bool, string) {
	//Checar se existe email
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	condicion := bson.M{"email": email} //consultar email en la base dados

	var result models.User
	err := col.FindOne(ctx, condicion).Decode(&result)
	ID := result.Id.Hex() //el resultado viene en exa
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
