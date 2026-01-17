package controller

import (
	"encoding/json"
	"go_movies/src/structs"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)

	for i, movie := range structs.Movies {
		if movie.ID == params["id"] {
			structs.Movies = append(structs.Movies[:i], structs.Movies[i+1:]...)

			var movie structs.Movie

			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			movie.ID = params["id"]

			structs.Movies[i] = movie

			if err := json.NewEncoder(w).Encode(movie); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}
