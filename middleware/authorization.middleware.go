package middleware

import (
	"fmt"
	"net/http"
	"time"

	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret_key")

func GenerateJWT(userName string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &m.Claims{
		Username:   userName,
		Authorized: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "",
			err
	}

	return tokenString, nil
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Hubo un error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}
			if token.Valid {
				endpoint(rw, r)
			}
		} else {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
