package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	h "github.com/AlexRojasB/go-mongoAtlas-connection.git/handlers"
	"github.com/AlexRojasB/go-mongoAtlas-connection.git/middleware"
)

func home(w http.ResponseWriter, R *http.Request) {
	fmt.Println("Home")
}
func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/refresh", h.RefreshToken)
	http.HandleFunc("/signup", h.SignUp)
	http.Handle("/feed", middleware.IsAuthorized(h.GetAllNewFeed))
	http.Handle("/ownfeed", middleware.IsAuthorized(h.GetAllMyFeeds))
	//http.Handle("/home", middleware.IsAuthorized(h.SignUp))
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	fmt.Println("API Started")
	handleRequests()
}
