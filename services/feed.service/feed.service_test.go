package feed_service_test

import (
	"testing"
	"time"

	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	feedService "github.com/AlexRojasB/go-mongoAtlas-connection.git/services/feed.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var feedId = "6169b4dcbdac5eb7644c1d30"

func TestCreateFeed(t *testing.T) {
	old, _ := primitive.ObjectIDFromHex(feedId)
	ownerId, _ := primitive.ObjectIDFromHex("616883bba8f0a40adbd10137")

	feed := m.Feed{
		ID:      old,
		Picture: "https://cdn.pixabay.com/photo/2015/06/19/09/39/lonely-814631_960_720.jpg",
		Likes:   55,
		Owner: m.Owner{
			ID:       ownerId,
			Name:     "Alexander Rojas",
			Location: "Santa Fe, Quesada",
			Picture:  "https://cdn.pixabay.com/photo/2017/02/08/16/45/man-2049447_960_720.jpg",
		},
		UpdatedAt: time.Now(),
	}
	id, err := feedService.Create(feed)

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	} else {
		t.Logf("La prueba finalizo con exito new ID: %s", id)
	}
}

func TestReadAllFeeds(t *testing.T) {
	skipAmount := 10

	feeds, err := feedService.Read(skipAmount)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	} else if feeds == nil || len(feeds) <= 0 {
		t.Error(err.Error())
		t.Fail()
	} else {
		t.Log("La prueba finalizo correctamente")
	}
}

func TestReadAllFeedsFromUser(t *testing.T) {
	ownerId := "616883bba8f0a40adbd10137"

	feeds, err := feedService.ReadFromOwner(ownerId)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	} else if feeds == nil || len(feeds) <= 0 {
		t.Error(err.Error())
		t.Fail()
	} else {
		t.Log("La prueba finalizo correctamente")
	}
}

func TestUpdateCommentsOrLike(t *testing.T) {
	old, _ := primitive.ObjectIDFromHex(feedId)
	ownerId, _ := primitive.ObjectIDFromHex("616896d2aab7533815351c51")
	var comments m.Comments
	comment := m.Comment{
		ID:      primitive.NewObjectID(),
		Comment: "This is a test comment",
		Owner: m.Owner{
			ID:       ownerId,
			Name:     "Alexander Benavides",
			Location: "Magallanes, San Ramon",
			Picture:  "https://cdn.pixabay.com/photo/2017/02/08/16/45/man-2049447_960_720.jpg",
		},
	}
	comments = append(comments, &comment)

	feed := m.Feed{
		ID:        old,
		Likes:     56,
		Comments:  comments,
		UpdatedAt: time.Now(),
	}

	err := feedService.Update(feed)
	if err != nil {
		t.Error("Se ha presentado un error")
		t.Fail()
	}
	t.Log("La prueba finalizo correctamente")
}

func TestDelete(t *testing.T) {
	err := feedService.Delete(feedId)
	if err != nil {
		t.Error("Se ha presentado un error")
		t.Fail()
	}
	t.Log("La prueba finalizo correctamente")
}
