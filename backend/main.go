package main

import (
	"fmt"
	"log"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	rt "github.com/Foodut/backend/router"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to Database
	_ = dbController.GetConnection()

	// Automigrate model to database
	// dbController.Migrate(db)

	router := mux.NewRouter()

	rt.LoadRouter(router)

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
