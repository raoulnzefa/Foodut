package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllCustomerWithAssociation(writer http.ResponseWriter, req *http.Request) {

	// Check store_name query
	customerId := req.URL.Query()["user_id"]

	// Get list of user object
	var customer []model.Customer = srvc.SearchCustByIdWithAssociation(customerId)

	// Set response
	var response rspn.Response
	if len(customer) > 0 {
		response.Response_200(customer)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func GetAllCustomerWithAssociationComplete(writer http.ResponseWriter, req *http.Request) {

	// Check store_name query
	customerId := req.URL.Query()["user_id"]

	// Get list of user object
	var users []dto.GetCustomer = srvc.SearchCustByIdWithAssociationComplete(customerId)

	// Set response
	var response rspn.Response
	if len(users) > 0 {
		response.Response_200(users)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func PostCustomer(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postUserDto dto.PostUser
	err := json.NewDecoder(req.Body).Decode(&postUserDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.MapToCustomer(postUserDto)

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
