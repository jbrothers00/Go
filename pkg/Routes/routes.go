package routes

import (
	"github.com/gorilla/mux"
	"github.com/jbrothers0028/Go/pkg/Controller"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", Controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", Controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", Controller.GetBooksById).Methods("GET")
	router.HandleFunc("/book/{id}", Controller.UpdateBooksById).Methods("PUT")
	router.HandleFunc("/book/{id}", Controller.DeleteBooksById).Methods("DELETE")
}
