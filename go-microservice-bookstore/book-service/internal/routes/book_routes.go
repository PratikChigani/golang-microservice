package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pratikchigani/book-service/internal/handlers"
	"github.com/pratikchigani/book-service/internal/middlewares"
)

func RegisterBookRoutes(r *mux.Router) {
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")

	r.Handle("/books", middleware.EnsureJSONMiddleware(http.HandlerFunc(handlers.CreateBook))).Methods("POST")
	r.Handle("/books/{id}", middleware.EnsureJSONMiddleware(http.HandlerFunc(handlers.UpdateBook))).Methods("PUT")

	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
}
