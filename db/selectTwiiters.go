package db

import (
	"context"
	"log"
	"time"

	"redsocial/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Twitters Paginated
func SelectTwiiters(ID string, page int64) ([]*models.ReturnTwitterRequest, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	dbase := MongoCN.Database("twittor")
	collect := dbase.Collection("tweets")

	var results []*models.ReturnTwitterRequest

	condition := bson.M{
		"iserid": ID,
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}}) //Ordenar por fecha Value -1 = descendente
	options.SetSkip((page - 1) * 20)                  //La primer vez que entra no saltea por 20, y despues paginando

	cursor, err := collect.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	//Percorrer los documentos por cursor
	//TODO cria un contexto vaxio
	for cursor.Next(context.TODO()) {
		var record models.ReturnTwitterRequest
		err := cursor.Decode(&record)

		if err != nil {
			return results, false
		}
		results = append(results, &record) //agregar en un slide un elemento
	}

	return results, true
}
