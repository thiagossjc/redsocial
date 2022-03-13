package db

import (
	"context"
	"time"

	"redsocial/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Graba el Tweet en el Banco de Dados
func InsertTweet(t models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbase := MongoCN.Database("twittor")
	collect := dbase.Collection("tweets")

	register := bson.M{
		"userID":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := collect.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	//extray la clave del ultimo registro insertado en el banco
	objId, _ := result.InsertedID.(primitive.ObjectID) //Buscando elID da inserção en el object

	return objId.String(), true, nil
}
