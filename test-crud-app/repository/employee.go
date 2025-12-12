package repository

import (
	"context"

	"github.com/kushvardhan/GO-Projects/model"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type EmployeeRepo struct{
	MongoCollection *mongo.Collection
}


func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (interface{}, error){
	result , err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil{
		return nil, err
	}

	return result, nil
}

func (r *EmployeeRepo) FindEmployeeById(empId string) (*model.Employee, error){
	var emp model.Employee

	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key:"employee_id", Value: empId}}).Decode(&emp)

	if err != nil{
		return nil, err
	}

	return &emp, nil
}

func 