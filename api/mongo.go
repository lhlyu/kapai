package api

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var collection *mongo.Collection

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Panic(err)
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		log.Panic(err)
	}
	collection = client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COL"))
}
