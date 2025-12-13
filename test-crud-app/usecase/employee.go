package usecase

import "net/http"

type EmployeeService struct{
	MongoCollection *mongo.Collection
}

type Response struct{
	Data interface{} `json:"data.omitempty"`
	Error string `json:"error, omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request){

}

func (svc *EmployeeService) GetEmployeeById(w http.ResponseWriter, r *http.Request){

}
func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request){

}
func (svc *EmployeeService) UpdateEmployeeById(w http.ResponseWriter, r *http.Request){

}
func (svc *EmployeeService) DeleteEmployeeById(w http.ResponseWriter, r *http.Request){

}

func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request){

}