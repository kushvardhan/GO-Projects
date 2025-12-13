package repository

import (
	"context"
	"fmt"

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

	return result.InsertedID, nil
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

func (r *EmployeeRepo) FindAllEmployee() ([]model.Employee, error){
	result , err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []model.Employee
	err = result.All(context.Background(),&emps)
	if err != nil{
		return nil, fmt.Errorf("Results decode error %s", err.Error())
	}
	return emps, nil
}

func (r *EmployeeRepo) UpdateEmployeeById(empId string, updateEmp *model.Employee) (int64, error){
	res, err:= r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key:"employee_Id",Value:empId}},
		bson.D{{Key:"$set", Value:updateEmp}})
	
	if err != nil{
		return 0, err
	}
	return res.ModifiedCount,nil
}	

func (r *EmployeeRepo) DeleteEmployeeById(empId string) (int64, error){
	res, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key:"employee_Id",Value:empId}})

	if err != nil{
		return 0,err
	}
	
	return res.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmployee() (int64, error){
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D())

	if err != nil{
		return 0, err
	}

	return result.DeletedCount, nil
}