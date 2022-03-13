package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Borro Tweet borra un twitter determinado
func DeleteTwitter(ID string, UserID string) error {
	ctx, cancel := (context.WithTimeout(context.Background(), time.Second*15))
	defer cancel()

	dbase := MongoCN.Database("twittor")
	collect := dbase.Collection("tweets")
	objId, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":     objId,
		"user_id": UserID,
	}
	_, err := collect.DeleteOne(ctx, condition)
	return err
}
