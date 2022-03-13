package controllers

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/models"
	service "github.com/Foodut/backend/services"
)

// func GetAllProducts(w http.ResponseWriter, r *http.Request) {
// 	db := GetConnection()

// 	var products []models.Product

// 	name := r.URL.Query()["product_name"]
// 	if name != nil {
// 		db.Where("product_name = ?", name[0]).First(&products)
// 	} else {
// 		db.Find(&products)
// 	}

// 	// Set response
// 	var response models.ProductResponse
// 	if len(products) > 0 {
// 		response.Status = 200
// 		response.Message = "Success Get Product"
// 		response.Data = products
// 	} else {
// 		response.Status = 204
// 		response.Message = "Not Found, No Content"
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func GetAllProducts(writer http.ResponseWriter, req *http.Request) {
	nameToFind := req.URL.Query()["product_name"]

	// Get product and check by query product_name
	var products []model.Product = service.SearchByName(nameToFind)

	// Set response
	var response model.ProductResponse
	if len(products) > 0 {
		// createDataResponse(products, response)
		response.Status = 200
		response.Message = "Success Get Product"
		response.Data = products
	} else {
		// createEmptyResponse(response)
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
