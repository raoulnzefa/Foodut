package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	prController "github.com/Foodut/backend/modules/product/rest-api/controller"
	trController "github.com/Foodut/backend/modules/transaction/rest-api/controller"
	usrController "github.com/Foodut/backend/modules/user/rest-api/controller"

	dbController "github.com/Foodut/backend/database"
)

func main() {
	// Connect to Database
	db := dbController.GetConnection()

	// Automigrate model to database
	dbController.Migrate(db)

	router := mux.NewRouter()

	router.HandleFunc("/usersTest", usrController.GetUserResponse).Methods("GET")

	//  USER
	//
	//- Get All User
	router.HandleFunc("/users", usrController.GetAllUsers).Methods("GET")

	//  PRODUCT
	//
	//- Get All Product
	router.HandleFunc("/products", prController.GetAllProducts).Methods("GET")

	//  TRANSACTION
	//
	//- Get All Transaction
	router.HandleFunc("/transactions", trController.GetAllTransactions).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
