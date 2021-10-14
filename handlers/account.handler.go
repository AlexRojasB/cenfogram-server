package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	middleware "github.com/AlexRojasB/go-mongoAtlas-connection.git/middleware"
	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	userService "github.com/AlexRojasB/go-mongoAtlas-connection.git/services/user.service"
	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials m.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//expectedPass, ok := users[credentials.Username]
	//if !ok || expectedPass != credentials.Password {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	loggedUser, err := userService.Read(credentials)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tokenString, err := middleware.GenerateJWT(credentials.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Token", tokenString)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loggedUser)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user m.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = userService.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenString, err := middleware.GenerateJWT(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Token", tokenString)
	w.WriteHeader(http.StatusCreated)
}

// not important
func ForgotPassword(w http.ResponseWriter, R *http.Request) {

}

// not important
func DeleteAccount(w http.ResponseWriter, r *http.Request) {

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
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
