// main.go
package main

import (
	"log"
	"net/http"

	"github.com/pratikchigani/user-service/internal/db"
	"github.com/pratikchigani/user-service/internal/routes"
)

func main() {
	db.InitDB() // opens SQLite DB connection
	log.Println("Starting user-service on port 9001...")
	log.Fatal(http.ListenAndServe(":9001", routes.RegisterRoutes()))
}
