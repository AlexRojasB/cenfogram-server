package handlers

import (
	"encoding/json"
	"net/http"

	middleware "github.com/AlexRojasB/go-mongoAtlas-connection.git/middleware"
	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	userService "github.com/AlexRojasB/go-mongoAtlas-connection.git/services/user.service"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials m.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

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

	newId, err := userService.Create(user)

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newId)
	w.WriteHeader(http.StatusCreated)
}

// not important
func ForgotPassword(w http.ResponseWriter, R *http.Request) {

}

// not important
func DeleteAccount(w http.ResponseWriter, r *http.Request) {

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var user m.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
