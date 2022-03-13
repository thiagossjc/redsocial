package db

import (
	"context"
	"time"

	"github.com/thiagossjc/redsocial/models"
)

func DeleteRelationship(Relat models.Relationship) (bool, error) {
	ctx, cancel := (context.WithTimeout(context.Background(), time.Second*15))
	defer cancel()

	dbase := MongoCN.Database("twittor")
	collect := dbase.Collection("relations")

	_, err := collect.DeleteOne(ctx, Relat)

	if err != nil {
		return false, err
	}
	return true, nil

}
