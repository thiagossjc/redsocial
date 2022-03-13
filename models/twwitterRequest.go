package models

//Captura del Body, el mensage que nos llega
type TwitterRequest struct {
	Message string `bson:"message" json:"message"`
}
