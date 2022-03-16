package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllSeller(w http.ResponseWriter, r *http.Request) {

	// Get list of seller object
	var sellers []dto.SellerMinimal = srvc.EmptySellerMinimal()

	// Set response
	var response rspn.Response
	if len(sellers) > 0 {
		response.Response_200(sellers)
	} else {
		response.Response_204()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetSellerByStoreWithProducts(writer http.ResponseWriter, req *http.Request) {

	// Check product_name query
	storeName := req.URL.Query()["store_name"]

	// Get list of seller object
	var seller model.Seller = srvc.SearchByStoreName(storeName)

	// Set response
	var response rspn.Response
	if seller.UserID != 0 {
		response.Response_200(seller)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
