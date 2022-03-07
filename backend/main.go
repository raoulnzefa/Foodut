package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/Foodut/backend/controllers"
)

func main() {
	controller.Connect()
	// db := controller.Connect()
	// db.AutoMigrate(&model.Admin{}, &model.Category{}, &model.Customer{}, &model.DetailProduct{}, &model.Product{}, &model.Seller{}, &model.Transaction{}, &model.User{})
	router := mux.NewRouter()

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
