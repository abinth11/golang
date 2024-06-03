package movieController

import (
	"encoding/json"
	"net/http"

	movieHelpers "github.com/abinth11/gowithmongodb/helpers"
	movies "github.com/abinth11/gowithmongodb/models"
	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	movies := movieHelpers.FindAll()
	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	movie, _ := movieHelpers.FindOne(params["movieId"])
	json.NewEncoder(w).Encode(movie)
}

func SaveMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie movies.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movieHelpers.InsertOne(movie)
	json.NewEncoder(w).Encode("Saved new movie")
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	movieHelpers.DeleteOne(params["movieId"])
	json.NewEncoder(w).Encode("Deleted movie")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	movieHelpers.DeleteAll()
	json.NewEncoder(w).Encode("Deleted all movies")
}

func UpdateMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	movieId := mux.Vars(r)["movieId"]
	var movie movies.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movieHelpers.UpdateOne(movieId, movie)
	json.NewEncoder(w).Encode("Updated the movie")
}
