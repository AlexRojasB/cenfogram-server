package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feed struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Picture  string             `json:"picture"`
	Likes    int32              `json:"likes"`
	Comments Comments           `json:"comments"`
	Owner    Owner              `json:"owner"`
}

type Feeds []*Feed
