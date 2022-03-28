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
