package db

import (
	"context"
	"time"

	"github.com/thiagossjc/redsocial/models"
)

//Graba la relación entre los usuários
func InsertRelationship(relat models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //agregar contexto miniatura para el tem de timeout
	defer cancel()

	dbase := MongoCN.Database("twittor")
	collect := dbase.Collection("relations")

	_, err := collect.InsertOne(ctx, relat)

	if err != nil {
		return false, err
	}

	return true, nil

}
