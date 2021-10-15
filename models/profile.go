package models

import (
	"time"
)

type Profile struct {
	Name     string    `json:"name"`
	Picture  string    `json:"picture"`
	Birthday time.Time `json:"birthday"`
}

type Profiles []*Profile
