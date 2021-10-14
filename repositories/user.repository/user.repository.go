package user_repository

import (
	"context"
	"time"

	"github.com/AlexRojasB/go-mongoAtlas-connection.git/database"
	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(user m.User) error {

	var err error

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Users, error) {
	var users m.Users
	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user m.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func Update(user m.User, userId string) error {
	var err error

	old, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": old}
	update := bson.M{"$set": bson.M{
		"name":       user.Name,
		"email":      user.Email,
		"updated_at": time.Now(),
	},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {
	var err error
	var old primitive.ObjectID

	old, err = primitive.ObjectIDFromHex(userId)
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
