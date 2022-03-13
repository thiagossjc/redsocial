package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MOdels User
type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	LastName  string             `bson:"last_name,omitempty" json:"lastName"`
	DateBirth time.Time          `bson:"date_birth,omitempty" json:"dateBirth"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	WebSite   string             `bson:"web_site" json:"webSite,omitempty"`
}
