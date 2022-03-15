package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/product/domain/model"
	service "github.com/Foodut/backend/modules/product/domain/service"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllProducts(writer http.ResponseWriter, req *http.Request) {

	// Get product and check by query
	var products []model.Product = service.EmptySearchBy()

	// Set response
	var response rspn.Response
	if len(products) > 0 {
		response.Response_200(products)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func GetProductByName(writer http.ResponseWriter, req *http.Request) {

	// Check product_name query
	nameToFind := req.URL.Query()["product_name"]

	// Get product and check by query
	var products []model.Product = service.SearchByName(nameToFind)

	// Set response
	var response rspn.Response
	if len(products) > 0 {
		response.Response_200(products)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
