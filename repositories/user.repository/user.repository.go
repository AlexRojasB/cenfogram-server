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

func Create(user m.User) (interface{}, error) {
	var dbUser m.User
	var err error

	filter := bson.M{"email": user.Email}

	cur := collection.FindOne(ctx, filter)
	errFind := cur.Decode(&dbUser)
	if errFind == nil {
		return nil, errors.New("The user already exists")
	}

	insertedUser, err := collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}
	return insertedUser.InsertedID, nil
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
