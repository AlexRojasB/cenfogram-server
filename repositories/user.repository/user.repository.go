package user_repository

import (
	"context"
	"errors"
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
	var users m.Users

	filter := bson.M{"email": user.Email}

	cur, er := collection.Find(ctx, filter)
	if er != nil {
		return er
	}

	for cur.Next(ctx) {
		var dbUser m.User
		err = cur.Decode(&dbUser)
		if err != nil {
			return err
		}
		users = append(users, &user)
	}

	if users != nil && len(users) > 0 {
		return errors.New("The user already exists")
	}

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

func Read(loginUser m.User) (m.User, error) {
	var user m.User
	filter := bson.M{"email": loginUser.Email, "password": loginUser.Password}

	cur := collection.FindOne(ctx, filter)
	err := cur.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func Update(user m.User, userId string) error {
	var err error

	old, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": old}
	update := bson.M{"$set": bson.M{
		"nick":       user.Nick,
		"password":   user.Password,
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
