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
	controller.Connect()
	db := controller.Connect()
	db.Debug().AutoMigrate(&model.User{}, &model.Admin{} /* &model.Category{}, &model.Customer{}, &model.DetailProduct{}, &model.Product{}, &model.Seller{}, &model.Transaction{}, ,*/)
	router := mux.NewRouter()

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
