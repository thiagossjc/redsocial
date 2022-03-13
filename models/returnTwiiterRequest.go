package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnTwitterRequest struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId  string             `bson:"user_id" json:"userID,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
