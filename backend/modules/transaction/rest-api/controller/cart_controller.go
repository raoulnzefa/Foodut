package controller

import (
	"encoding/json"
	"net/http"

	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
	dto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func PostToCart(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postCartDto dto.PostCart
	err := json.NewDecoder(req.Body).Decode(&postCartDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendCartForCreate(postCartDto)

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

func UpdateCart(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postCartDto dto.PostCart
	err := json.NewDecoder(req.Body).Decode(&postCartDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendCartForUpdate(postCartDto)

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

func DeleteSpesificCart(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var delSpCart dto.DeleteSpecificCart
	err := json.NewDecoder(req.Body).Decode(&delSpCart)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendCartForDelSpesific(delSpCart)

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
