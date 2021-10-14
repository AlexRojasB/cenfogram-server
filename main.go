package main

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/AlexRojasB/go-mongoAtlas-connection.git/handlers"
	middleware "github.com/AlexRojasB/go-mongoAtlas-connection.git/middleware"
	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	userService "github.com/AlexRojasB/go-mongoAtlas-connection.git/repositories/user.repository"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	user := m.User{
		Name:  "Alejandrina",
		Email: "alejandrinamirandagarcia@gmail.com",
	}
	err := userService.Create(user)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequests() {
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/refresh", h.Refresh)
	http.Handle("/home", middleware.IsAuthorized(h.Home))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("API Started")
	handleRequests()
}
