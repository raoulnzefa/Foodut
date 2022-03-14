package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/Foodut/backend/controllers"
	dbController "github.com/Foodut/backend/database"
)

func main() {
	// Connect to Database
	db := dbController.GetConnection()

	// Automigrate model to database
	dbController.Migrate(db)

	router := mux.NewRouter()

	router.HandleFunc("/usersTest", controller.GetUserResponse).Methods("GET")

	//  USER
	//
	//- Get All User
	router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")

	//  PRODUCT
	//
	//- Get All Product
	router.HandleFunc("/products", controller.GetAllProducts).Methods("GET")

	//  TRANSACTION
	//
	//- Get All Transaction
	router.HandleFunc("/transactions", controller.GetAllTransactions).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
