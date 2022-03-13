package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Twitters de los Usuarios que seguimos
type TwittersFollowers struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID           string             `bson:"user_id" json:"userId,omitempty"`
	UserRelationship string             `bson:"user_relationship" json:"userRelationship,omitempty"`
	Twwiter          struct {
		Messaage string    `bson:"message" json:"message,omitempty"`
		Date     time.Time `bson:"date" json:"date,omitempty"`
		ID       string    `bson:"_id" json:"_id,omitempty"`
	}
}
