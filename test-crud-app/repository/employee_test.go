package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/kushvardhan/GO-Projects/model"
	"go.mongodb.org/mongo-driver/internal/uuid"
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

func TestMongoOperations(t *testing.T){
	mongoTestClient := NewMongoClient();
	defer mongoTestClient.Disconnect(context.Background())

	emp1:= uuid.New().String()
	// emp2:= uuid.New().String()

	call := mongoTestClient.Database("test-app").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: call}

	t.Run("Insert Employee 1", func(t *testing.T){
		emp:= model.Employee{
			Name :"Tony Stark",
			Department: "Physics",
			EmployeeId : emp1,
		}

		result, err := empRepo.InsertEmployee(&emp);

		if err != nil{
			t.Fatal("Insert 1 operation failed", err)
		}

		t.log("Insert 1 successfull", result)

	})

	t.Run("Get Employee 1", func(t *testing.T){
		result , err := empRepo.FindEmployeeById(emp1)

		if err != nil{
			t.Fatal("get operation failed", err)
		}
		t.Log("emp 1", result.Name)
	})

	t.Run("Update Employee 1 Name", func(t *testing.T){
		emp := model.Employee{
			Name :"Tony Stark aka Iron Man",
			Department: "Physics",
			EmployeeId: "emp1",
		}

		result, err := empRepo.UpdateEmployeeById(emp1, &emp)
		if err != nil{
			log.Fatal("Error while updating name", err);
		}

		t.Log("emp 1", result.Name);
	})

	

}