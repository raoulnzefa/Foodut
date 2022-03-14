package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/product/domain/model"
	service "github.com/Foodut/backend/modules/product/domain/service"
)

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
