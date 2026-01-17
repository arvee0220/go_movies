package controller

import (
	"encoding/json"
	"go_movies/src/structs"
	"math/rand"
	"net/http"
	"strconv"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie structs.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	structs.Movies = append(structs.Movies, movie)

	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
