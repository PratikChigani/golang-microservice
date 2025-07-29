// routes/routes.go
package routes

import (
	"net/http"

	"github.com/pratikchigani/user-service/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	return r
}
