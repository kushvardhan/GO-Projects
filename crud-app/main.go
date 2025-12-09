package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Tite string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	LastName string `json:"lastname"`
	FirstName string `json:"firstname"`
}

var movies[]Movie

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/movies", handleGetMovies);
	r.HandleFunc("")
	
	fmt.Fprintf("Server Starting");
}