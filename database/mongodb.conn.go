package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "seke"
	pwd      = "epsi1234"
	host     = "localhost"
	port     = 27017
	database = "tutorial"
)

func GetCollection(collection string) *mongo.Collection {
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	//client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	//if err != nil {
	//	panic(err.Error())
	//}

	//ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	//err client.Connect(ctx)

	//if err != nil {
	//	panic(err.Error())
	//}
	clientOptions := options.Client().ApplyURI("mongodb+srv://seke:upsi1234@cluster0.zrmzp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(database).Collection(collection)
}
