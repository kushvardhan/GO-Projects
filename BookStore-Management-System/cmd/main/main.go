package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main(){
	r := mux.NewRouter();
	log.Println("Starting Server at PORT:8000.");
	log.Fatal(http.ListenAndServe("8000",r));
}