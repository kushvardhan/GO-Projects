package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client;

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load error", err)
	}

	log.Println("env file loaded")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL"))

	if err != nil {
		log.Fatal("Connection error", err)
	}

	err = mongoClient.Ping(context.Backgound(), readpref.Primary())
	if err != nil{
		log.Fatal("ping failed", err)
	}
	log.Println("mongo connected");
}

func main(){
	r := mux.NewRouter()
	log.Println("Starting Server at PORT:8000.");
	log.Fatal(http.ListenAndServe("PORT:8000", r));
}