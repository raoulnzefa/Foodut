package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/Foodut/backend/controllers"
	model "github.com/Foodut/backend/models"
)

func main() {
	db := controller.Connect()
	db.Debug().AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Customer{},
		&model.Seller{},
		&model.DetailProduct{},
		&model.Transaction{},
		&model.Category{},
		&model.Product{},
		&model.Picture{},
	)

	router := mux.NewRouter()

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
