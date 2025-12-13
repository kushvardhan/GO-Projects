package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kushvardhan/GO-Projects/usecase"
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

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL"));

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
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	empService := usecase.EmployeeService{MongoCollection: coll}

	r := mux.NewRouter()

	r.HandleFunc("/health", healtHandler).Methids(http.MethodGet)

	r.HandleFunc("/employee", empService.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee", empService.GetAllEmployee).Methods("GET");
	r.HandleFunc("/employee/{id}",empService.GetEmployeeById).Methods("POST");
	r.HandleFunc("/employee/{id}",empSevice.UpdateEmployeeById).Methods("PUT");
	r.HandleFunc("/employee/{id}", empService.DeleteEmployeeById).Methods("DELETE");
	r.HandleFunc("/employee",empService.DeleteAllEmployee).Methods("DELETE");


	log.Println("Sever is running on PORT: 4444");
	http.ListenAndServe(":4444", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([] byte{"running..."});
}