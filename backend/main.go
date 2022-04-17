package main

import (
	"fmt"
	"log"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	rt "github.com/Foodut/backend/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Connect to Database
	db := dbController.GetConnection()

	// Automigrate model to database
	dbController.Migrate(db)

	router := mux.NewRouter()

	rt.LoadRouter(router)

	// http.Handle("/", router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://127.0.0.1:8081", "http://localhost:1234", "http://127.0.0.1:1234"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", handler))
}
