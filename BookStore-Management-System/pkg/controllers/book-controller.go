package controllers

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kushvardhan/GO-Projects/BookStore-Management-System/pkg/utils"
	"github.com/kushvardhan/GO-Projects/BookStore-Management-System/pkg/models"
	"encoding/json"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Response){
	newBooks:=models.GetAllBooks()
	res,_ := json.Marshal(newBooks);
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}

func GetBookById(w http.ResponseWriter, r *http.Response){
	vars := mux.Vars(r);
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0);
	if err != nil{
		fmt.Println(err)
	}
	bookDetails,_ := models.GetBookById(Id);
	res,_ := json.Marshal(bookDetails);
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK);
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r,CreateBook)
	b:= CreateBook.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}