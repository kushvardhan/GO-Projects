package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(movies);
}

func deleteMovieById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json");
	params := mux.Vars(r);
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies);
}	

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies,movie);
	json.NewEncoder(w).Encode(movie);
}


func getMovieById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json");
	params := mux.Vars(r);

	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item);
			return;
		}
	}
}

func updateMovieById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for i, item := range movies {
        if item.ID == params["id"] {
            var updated Movie
            json.NewDecoder(r.Body).Decode(&updated)

            updated.ID = params["id"]
            movies[i] = updated

            json.NewEncoder(w).Encode(updated)
            return
        }
    }

    http.Error(w, "Movie not found", http.StatusNotFound)
}


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
	r.HandleFunc("/movies/{id}",getMovieById).Methods("GET");
	r.HandleFunc("/createMovie", createMovie).Methods("POST");
	r.HandleFunc("/movies/{id}",updateMovieById).Methods("PUT");
	r.HandleFunc("/movies/{id}",deleteMovieById).Methods("DELETE");

	
	fmt.Printf("Server Starting at PORT: 8000\n");
	log.Fatal(http.ListenAndServe(":8000",r))
}