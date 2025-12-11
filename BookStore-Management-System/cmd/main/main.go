package main

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kushvardhan/GO-Projects/BookStore-Management-System/pkg/routes"

)

func main(){
	r := mux.NewRouter();

	routes.RegisterBookStoreRoutes(r)
	http.Handle("r");

	log.Println("Starting Server at PORT:8000.");
	log.Fatal(http.ListenAndServe("localhost:8000",r));
}