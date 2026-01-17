package controller

import (
	"encoding/json"
	"go_movies/src/structs"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, movie := range structs.Movies {
		if movie.ID == params["id"] {
			structs.Movies = append(structs.Movies[:i], structs.Movies[i+1:]...)
			break
		}

		http.Error(w, "Movie not found", http.StatusNotFound)
	}

	if err := json.NewEncoder(w).Encode(structs.Movies); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
