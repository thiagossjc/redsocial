package db

import (
	"context"
	"fmt"
	"redsocial/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Select Trar Todos lo usuarios con relacion con el user logado
func SelectAllUsers(ID string, page int64, search string, typeR string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbase := MongoCN.Database("twittor")
	collect := dbase.Collection("users")

	var results []*models.User //sea un slice

	findoptions := options.Find()
	findoptions.SetSkip((page - 1) * 20)
	findoptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	curso, err := collect.Find(ctx, query, findoptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var found, included bool

	for curso.Next(ctx) {
		var user models.User
		err := curso.Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var relation models.Relationship
		relation.UserId = ID
		relation.UserRelationshipId = user.Id.Hex()
		included = false
		found, err = SelectRelationShip(relation)

		if typeR == "new" && !found {
			included = true

		}
		if typeR == "follow" && !found {
			included = true
		}
		if relation.UserRelationshipId == ID {
			included = false
		}
		if included == true {
			user.Password = ""
			user.Biography = ""
			user.WebSite = ""
			user.Location = ""
			user.Banner = ""
			user.Email = ""

			results = append(results, &user)
		}
	}
	err = curso.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	curso.Close(ctx)
	return results, true
}
