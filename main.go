package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmccerezo/movies-crud-api/models"
)

var movies []models.Movie

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies", GetAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovieByID).Methods("GET")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	movies = models.Movie{}.Init()

	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// * CREATE
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movie := models.Movie{}

	_ = json.NewDecoder(r.Body).Decode(&movie)

	for i := range movies {
		id, _ := strconv.Atoi(movies[i].ID)
		movie.ID = strconv.Itoa(id + 1)
	}

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movies)
}

// * READ
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, m := range movies {
		if m.ID == params["id"] {
			json.NewEncoder(w).Encode(m)
			return
		}
	}
}

// * UPDATE
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	movie := models.Movie{}

	for i, m := range movies {
		if m.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)

			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]
			movies = append(movies, movie)
		}
	}

	json.NewEncoder(w).Encode(movies)
}

// * DELETE
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, m := range movies {
		if m.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}
