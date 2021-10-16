package handlers

import (
	"encoding/json"
	"net/http"

	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	feedService "github.com/AlexRojasB/go-mongoAtlas-connection.git/services/feed.service"
)

func GetAllNewFeed(w http.ResponseWriter, r *http.Request) {
	feeds, err := feedService.Read(0)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(feeds)
}

func GetAllMyFeeds(w http.ResponseWriter, r *http.Request) {
	var credentials m.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	feeds, err := feedService.ReadFromOwner(credentials.ID.Hex())
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(feeds)
}

func PostNewFeed(w http.ResponseWriter, r *http.Request) {

}

func DeleteFeed(w http.ResponseWriter, r *http.Request) {

}

func PostCommentFeed(w http.ResponseWriter, r *http.Request) {

}

func PostLikeFeed(w http.ResponseWriter, r *http.Request) {

}
