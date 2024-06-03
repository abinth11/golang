package movieRoutes

import (
	movieController "github.com/abinth11/gowithmongodb/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies/all", movieController.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movies{movieId}", movieController.GetMovieById).Methods("GET")
	r.HandleFunc("/api/movies/save", movieController.SaveMovie).Methods("POST")
	r.HandleFunc("/api/movies/update/{movieId}", movieController.UpdateMovieById).Methods("PUT")
	r.HandleFunc("/api/movies/delete/all", movieController.DeleteAllMovies).Methods("DELETE")
	r.HandleFunc("/api/movies/delete/{movieId}", movieController.DeleteMovieById).Methods("DELETE")
	return r
}
