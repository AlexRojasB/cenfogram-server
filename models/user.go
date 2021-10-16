package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nick     string             `json:"nick"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Profile  Profile            `json:"profile"`
}

type Users []*User
