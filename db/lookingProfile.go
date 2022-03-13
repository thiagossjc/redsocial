package db

import (
	"context"
	"fmt"
	"time"

	"redsocial/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LokkingProfile(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dBase := MongoCN.Database("twittor")
	collect := dBase.Collection("users")

	var profile models.User //profile es perfil

	objId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objId,
	}
	err := collect.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return profile, err
	}
	return profile, nil
}
