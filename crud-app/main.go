package main

import (
	"fmt"
	"log"
//	"encoding/json"
//	"math/rand"
	"net/http"
//	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	LastName string `json:"lastname"`
	FirstName string `json:"firstname"`
}

var movies[]Movie



func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{
	ID:    "1",
	Isbn:  "43877",
	Title: "Movie One",
	Director: &Director{
		FirstName: "Anurag",
		LastName:  "Kashyap",
	},
})

movies = append(movies, Movie{
	ID:    "2",
	Isbn:  "55231",
	Title: "The Silent River",
	Director: &Director{
		FirstName: "Raj",
		LastName:  "Singh",
	},
})

movies = append(movies, Movie{
	ID:    "3",
	Isbn:  "98412",
	Title: "Shadows of Time",
	Director: &Director{
		FirstName: "Neeraj",
		LastName:  "Ghosh",
	},
})

movies = append(movies, Movie{
	ID:    "4",
	Isbn:  "77129",
	Title: "Echoes of Destiny",
	Director: &Director{
		FirstName: "Reema",
		LastName:  "Chopra",
	},
})

movies = append(movies, Movie{
	ID:    "5",
	Isbn:  "12984",
	Title: "The Lost Path",
	Director: &Director{
		FirstName: "Vikram",
		LastName:  "Menon",
	},
})

movies = append(movies, Movie{
	ID:    "6",
	Isbn:  "33941",
	Title: "Dreamcatcher",
	Director: &Director{
		FirstName: "Suraaj",
		LastName:  "Verma",
	},
})

movies = append(movies, Movie{
	ID:    "7",
	Isbn:  "44891",
	Title: "Neon Nights",
	Director: &Director{
		FirstName: "Farhan",
		LastName:  "Ali",
	},
})

movies = append(movies, Movie{
	ID:    "8",
	Isbn:  "99321",
	Title: "Beyond Reality",
	Director: &Director{
		FirstName: "Kiran",
		LastName:  "Shah",
	},
})

movies = append(movies, Movie{
	ID:    "9",
	Isbn:  "66541",
	Title: "A Journey Unknown",
	Director: &Director{
		FirstName: "Amrita",
		LastName:  "Kapoor",
	},
})


	r.HandleFunc("/movies", getMovies).Methods("GET");
	r.HandleFunc("/movied/{id}",getMovieById).Methods("GET");
	r.HandleFunc("/createMovie", createMovies).Methods("POST");
	r.HandleFunc("/movied/{id}",updateMovieById).Methods("PUT");
	r.HandleFunc("/movied/{id}",deleteMovieById).Methods("DELETE");

	
	fmt.Printf("Server Starting at PORT: 8000\n");
	log.Fatal(http.ListenAndServe(":8000",r))
}