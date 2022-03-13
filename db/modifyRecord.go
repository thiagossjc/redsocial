package db

import (
	"context"
	"time"

	"redsocial/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Update register de Usuer
func ModifyRecord(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	dbas := MongoCN.Database("twittor")
	collect := dbas.Collection("users")

	record := make(map[string]interface{}) //cria un registro vacio de tipo interface
	if len(u.Name) > 0 {                   //Lhe grava el valor que haga falta lhe pasando los parametros
		record["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		record["last_name"] = u.LastName
	}

	record["date_birth"] = u.DateBirth

	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}
	if len(u.Location) > 0 {
		record["location"] = u.Location
	}
	if len(u.Biography) > 0 {
		record["biography"] = u.Banner
	}
	if len(u.WebSite) > 0 {
		record["web_site"] = u.WebSite
	}

	updateString := bson.M{
		"$set": record, //quando quiero atualizar un registro en Mongo tengo que poner eso
	}

	objID, _ := primitive.ObjectIDFromHex(ID)     //converter nuesstro ID en ObjectID
	filter := bson.M{"_id": bson.M{"$eq": objID}} //"$eq" = equal filtro
	_, err := collect.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
