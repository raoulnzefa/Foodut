package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllSeller(writer http.ResponseWriter, req *http.Request) {

	// Get list of seller object
	var sellers []dto.SellerMinimal = srvc.EmptySellerMinimal()

	// Set response
	var response rspn.Response
	if len(sellers) > 0 {
		response.Response_200(sellers)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func GetSellerByStoreWithProducts(writer http.ResponseWriter, req *http.Request) {

	// Check store_name query
	storeName := req.URL.Query()["store_name"]

	// Get list of seller object
	var seller model.Seller = srvc.SearchByStoreName(storeName)

	// Set response
	var response rspn.Response
	if seller.UserID > 0 {
		response.Response_200(seller)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func PostSeller(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postSellerDto dto.PostSeller
	err := json.NewDecoder(req.Body).Decode(&postSellerDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.MapToSeller(postSellerDto)

	// Set response
	var response rspn.Response
	if result.Error == nil {
		response.Response_201()
	} else {
		response.Response_400(result.Error)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
