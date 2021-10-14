package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	Location  string             `json:"location"`
	Picture   string             `json:"picture"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Owners []*Owner
