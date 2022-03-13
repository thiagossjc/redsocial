package models

type Relationship struct {
	UserId             string `bson:"user_id" json:"userID"`                     //mi id
	UserRelationshipId string `bson:"user_relationship" json:"userRelationship"` //id del user que estoy seguindo
}
