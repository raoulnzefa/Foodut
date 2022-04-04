package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/transaction/domain/model"
	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
	dto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func GetCartWithAvailability(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postTransactionDto dto.PostTransaction
	err := json.NewDecoder(req.Body).Decode(&postTransactionDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendCartForReadCartAvailability(postTransactionDto.CustomerId)

	// Set response
	var response rspn.Response
	if len(result) > 0 {
		response.Response_200(result)
	} else {
		response.Response_204("Get cart fail || Empty cart")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

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
		response.Response_201("Success post cart")
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
		response.Response_201("Success update cart")
	} else {
		response.Response_400(result.Error)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func DeleteSpesificProductFromCart(writer http.ResponseWriter, req *http.Request) {

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
		response.Response_201("Delete spesific cart success")
	} else {
		response.Response_400(result.Error)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func DeleteCarts(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var delCart []model.Cart
	err := json.NewDecoder(req.Body).Decode(&delCart)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendCartForDelete(delCart)

	// Set response
	var response rspn.Response
	if result.Error == nil {
		response.Response_201("Delete carts success")
	} else {
		response.Response_400(result.Error)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
