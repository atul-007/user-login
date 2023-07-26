package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName    string             `json:"username,omitempty"`
	Password    string             `json:"password,omitempty"`
	Dob         string             `json:"dob"`
	Address     string             `json:"address"`
	Description string             `json:"description"`
}
