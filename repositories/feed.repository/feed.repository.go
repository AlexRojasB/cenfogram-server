package feed_repository

import (
	"context"
	"errors"
	"time"

	"github.com/AlexRojasB/go-mongoAtlas-connection.git/database"
	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("feeds")
var ctx = context.Background()

func Create(feed m.Feed) (interface{}, error) {
	var dbFeed m.Feed
	var err error

	filter := bson.M{"_id": feed.ID}

	cur := collection.FindOne(ctx, filter)
	errFind := cur.Decode(&dbFeed)
	if errFind == nil {
		return nil, errors.New("The post already exists")
	}

	insertedFeed, err := collection.InsertOne(ctx, feed)

	if err != nil {
		return nil, err
	}
	return insertedFeed.InsertedID, nil
}

func Read(skipAmount int) (m.Feeds, error) {
	var feeds m.Feeds
	//filter := bson.M{"owner._id":userId}
	//jumpPagePipeline := bson.D{{"$skip", skipAmount}}
	//limitPipeline := bson.D{{"$limit", 10}}
	//var feedPipeline mongo.Pipeline
	//feedPipeline = append(feedPipeline, jumpPagePipeline)
	//feedPipeline = append(feedPipeline, limitPipeline)

	//cur, er := collection.Aggregate(ctx, feedPipeline)
	cur, er := collection.Find(ctx, bson.M{})
	if er != nil {
		return nil, er
	}

	for cur.Next(ctx) {
		var dbFeed m.Feed
		err := cur.Decode(&dbFeed)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, &dbFeed)
	}
	return feeds, nil
}

func ReadAllFromOwner(ownerId string) (m.Feeds, error) {
	var feeds m.Feeds
	ownerID, _ := primitive.ObjectIDFromHex(ownerId)
	filter := bson.M{"owner._id": ownerID}
	//jumpPagePipeline := bson.D{{"$skip", skipAmount}}
	//limitPipeline := bson.D{{"$limit", 10}}
	//var feedPipeline mongo.Pipeline
	//feedPipeline = append(feedPipeline, jumpPagePipeline)
	//feedPipeline = append(feedPipeline, limitPipeline)

	//cur, er := collection.Aggregate(ctx, feedPipeline)
	cur, er := collection.Find(ctx, filter)
	if er != nil {
		return nil, er
	}

	for cur.Next(ctx) {
		var dbFeed m.Feed
		err := cur.Decode(&dbFeed)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, &dbFeed)
	}
	return feeds, nil
}

func UpdateCommentsOrLikes(feed m.Feed) error {
	var err error

	filter := bson.M{"_id": feed.ID}
	update := bson.M{"$set": bson.M{
		"comments":   feed.Comments,
		"likes":      feed.Likes,
		"updated_at": time.Now(),
	},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(feedId string) error {
	var err error
	var old primitive.ObjectID

	old, err = primitive.ObjectIDFromHex(feedId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": old}
	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	return nil
}
