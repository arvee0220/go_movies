package controller

import (
	"encoding/json"
	"go_movies/src/structs"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(structs.Movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range structs.Movies {
		if movie.ID == params["id"] {
			err := json.NewEncoder(w).Encode(movie)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			break
		}
	}
}
