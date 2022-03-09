package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Foodut/backend/models"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	//defer db.Close()

	var products []models.Product

	name := r.URL.Query()["product_name"]
	if name != nil {
		db.Where("product_name = ?", name[0]).First(&products)
	} else {
		db.Find(&products)
	}

	// Set response
	var response models.ProductResponse
	if len(products) > 0 {
		response.Status = 200
		response.Message = "Success Get Product"
		response.Data = products
	} else {
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
