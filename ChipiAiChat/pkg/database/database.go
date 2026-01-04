package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo *mongo.Client

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:rootpass@localhost:27017"))
	if err != nil {
		log.Println(err)
	}

	Mongo = client
	log.Println("Connected to MongoDB")
}

func Messages() *mongo.Collection {
	return Mongo.Database("chipiai").Collection("messages")
}

func Rooms() *mongo.Collection {
	return Mongo.Database("chipiai").Collection("rooms")
}
