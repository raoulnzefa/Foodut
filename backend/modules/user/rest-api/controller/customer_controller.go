package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"

	"github.com/gorilla/mux"
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

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get id from path
	vars := mux.Vars(r)
	custID := vars["user_id"]

	var customer model.Customer
	// Get customer from database by id
	err = con.First(&customer, custID).Error

	// Get customer data
	// Belum bisa ambil nama customer nya soalnya beda table ada di parent
	name := r.Form.Get("name")
	fmt.Println("NAME : " + name)
	address := r.Form.Get("address")

	// Set inputted data to object
	if name != "" {
		customer.User.Name = name
	}
	if address != "" {
		customer.Address = address
	}

	var response rspn.Response
	if err == nil {
		result := con.Save(&customer)
		if result.Error == nil {
			response.Response_200(customer)
		} else {
			response.Response_400("Edit Customer Data Failed " + result.Error.Error())
		}
	} else {
		response.Response_400("Edit Customer Data Failed, ID Not Valid" + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
