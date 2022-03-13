package db

import (
	"context"
	"fmt"
	"redsocial/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//Busca Relationship
func SelectRelationShip(rel models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbase := MongoCN.Database("twittor")
	collect := dbase.Collection("relations")

	condition := bson.M{
		"user_id":           rel.UserId,
		"user_relationship": rel.UserRelationshipId,
	}
	var result models.Relationship
	fmt.Println(result)
	err := collect.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
