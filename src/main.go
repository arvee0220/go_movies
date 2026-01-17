package main

import (
	"fmt"
	"go_movies/src/controller"
	"go_movies/src/structs"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	structs.Movies = append(structs.Movies, Movie{ID: "1", Isbn: "438227", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	structs.Movies = append(structs.Movies, Movie{ID: "2", Isbn: "45455", Title: "Movie 2", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	r.HandleFunc("/movies", controller.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controller.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.DeleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
