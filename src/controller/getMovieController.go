package controller

import (
	"encoding/json"
	"go_movies/src/structs"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(structs.Movies); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Unable to fetch movies", http.StatusNotFound)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range structs.Movies {
		if movie.ID == params["id"] {

			if err := json.NewEncoder(w).Encode(movie); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			break
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}
