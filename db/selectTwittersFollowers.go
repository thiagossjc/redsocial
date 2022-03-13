package db

import (
	"context"
	"redsocial/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func SelectTwittersFollowers(ID string, page int) ([]models.TwittersFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	dbase := MongoCN.Database("twittor")
	collect := dbase.Collection("relations")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)

	conditions = append(conditions, bson.M{"$match": bson.M{"user_id": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",            //con que tabua queremos unir con relacion
			"localField":   "user_relationship", //campo que une una tabla con la otra
			"foreingField": "user_id",
			"as":           "tweets",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweets"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweets.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip}) //siempre skip antes del limite
	conditions = append(conditions, bson.M{"$limit": 20})
	cursor, err := collect.Aggregate(ctx, conditions)
	var results []models.TwittersFollowers
	err = cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}
	return results, true
}
