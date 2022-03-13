package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objecto e conexión a la base de dados*/
var MongoCN = ConectDB()

//aponta el camino del banco de dados mongodb
var clientOptions = options.Client().ApplyURI("mongodb+srv://thiagossjc:qvencedor07@cluster0.epyaj.mongodb.net/twittor?retryWrites=true&w=majority")

//Conexión con el Banco de Dado
func ConectDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) //Un ping en la base ddo

	if err != nil {
		log.Fatal(err) //function error converte el objetct en stri
		return client
	}
	log.Println("Cnexión Exitosa con la BD!")
	return client

}

//CheckConnection es el pingn la base de dados
func CheckConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
