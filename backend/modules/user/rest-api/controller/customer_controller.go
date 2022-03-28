package controller

import (
	"encoding/json"
	"net/http"


	dbController "github.com/Foodut/backend/database"
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
	var postCustDto dto.PostCustomer
	err := json.NewDecoder(req.Body).Decode(&postCustDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendForCreateCustomer(postCustDto)

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

func EditCustomer(w http.ResponseWriter, r *http.Request) {
	// Check connection
	con := dbController.GetConnection()

	// Get email from logged member
	email := GetEmailFromToken(r)

	var customer model.Customer
	con.Where("email=?", email).First(&customer)

	if name := r.Form.Get("name"); name != "" {
		customer.User.Name = name
	}
	if address := r.Form.Get("address"); address != "" {
		customer.Address = address
	}
	if password := r.Form.Get("password"); password != "" {
		customer.User.Password = password
	}

	result := con.Save(&customer)

	var response rspn.Response
	if result.Error == nil {
		response.Response_200("data has been updated")
	} else {
		response.Response_400(result)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
