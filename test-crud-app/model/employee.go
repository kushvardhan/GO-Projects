package model

type Employee struct{
	EmployeeId string `json:"employee_Id,omitempty" bson:"employee_Id"`
	Name string `json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`
}