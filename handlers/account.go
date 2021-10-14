package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	middleware "github.com/AlexRojasB/go-mongoAtlas-connection.git/middleware"
	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials m.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPass, ok := users[credentials.Username]
	if !ok || expectedPass != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := middleware.GenerateJWT(credentials.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Token", tokenString)
	//http.SetCookie(w, &http.Cookie{
	//Name:    "token",
	//Value:   tokenString,
	//Expires: expirationTime,
	//})
}

func Home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(fmt.Sprintf("Hello, %s", "antonio")))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value

	claims := &m.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return "jwtKey", nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString("jwtKey")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
