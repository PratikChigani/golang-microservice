package main

import (
	"log"
	"net/http"

	"github.com/pratikchigani/book-service/internal/db"
	"github.com/pratikchigani/book-service/internal/handlers"
	middleware "github.com/pratikchigani/book-service/internal/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	// Routes
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	log.Println("Book Service running on port 8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
