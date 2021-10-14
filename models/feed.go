package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feed struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Picture   string             `json:"picture"`
	Likes     int32              `json:"likes"`
	Comments  Comments           `json:"comments"`
	Owner     Owner              `json:"owner"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Feeds []*Feed
