package conf

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var  Mongoclient *mongo.Client

func init() {

	clientOptions := options.Client()
	mongoUrl := os.Getenv("mongoUrl")
	if mongoUrl == "" {
		log.Fatalf("mongo uri not defined in env")
	}
	uri := clientOptions.ApplyURI(mongoUrl)
	client, err := mongo.NewClient(uri)
	ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("connect mongodb failed : ",err)
	}
	log.Println("connect mongodb successful")
	Mongoclient=client
}