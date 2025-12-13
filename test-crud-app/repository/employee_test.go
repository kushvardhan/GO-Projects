package repository

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("MONGO_URL not set in environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal("error while connecting mongodb:", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("mongodb ping failed:", err)
	}

	log.Println("mongodb connected and ping successful")

	return client
}
