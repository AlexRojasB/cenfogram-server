package main

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/AlexRojasB/go-mongoAtlas-connection.git/handlers"
)

func home(w http.ResponseWriter, R *http.Request) {
	fmt.Println("Home")
}
func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/refresh", h.RefreshToken)
	http.HandleFunc("/signup", h.SignUp)
	//http.Handle("/home", middleware.IsAuthorized(h.SignUp))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("API Started")
	handleRequests()
}
