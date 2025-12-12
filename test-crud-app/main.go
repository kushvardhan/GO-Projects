package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	log.Println("Starting Server at PORT:8000.");
	log.Fatal(http.ListenAndServe("PORT:8000", r));
}