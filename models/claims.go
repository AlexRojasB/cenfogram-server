package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username   string `json:"username"`
	Authorized bool   `json:"authorized"`
	jwt.StandardClaims
}
