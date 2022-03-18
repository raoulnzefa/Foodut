package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	// Get list of user object
	var users []model.User = srvc.EmptyUserSearch()

	// Set response
	var response rspn.Response
	if len(users) > 0 {
		response.Response_200(users)
	} else {
		response.Response_204()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// TODO
// func GetUsersByUsernameLike(w http.ResponseWriter, r *http.Request) {

// 	// Get list of user object
// 	var users []model.User = srvc.EmptySearchBy()

// 	name := r.URL.Query()["name"]
// 	if name != nil {
// 		db.Where("name = ?", name[0]).First(&users)
// 	} else {
// 		db.Find(&users)
// 	}

// 	// Set response
// 	var response rspn.Response
// 	if len(users) > 0 {
// 		response.Response_200(users)
// 	} else {
// 		response.Response_204()
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

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
		fmt.Println(result.Error)
		response.Response_400()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func PostSeller(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postUserDto dto.PostUser
	err := json.NewDecoder(req.Body).Decode(&postUserDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.MapToSeller(postUserDto)

	// Set response
	var response rspn.Response
	if result.Error == nil {
		response.Response_201()
	} else {
		fmt.Println(result.Error)
		response.Response_400()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func PostAdmin(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postUserDto dto.PostUser
	err := json.NewDecoder(req.Body).Decode(&postUserDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.MapToAdmin(postUserDto)

	// Set response
	var response rspn.Response
	if result.Error == nil {
		response.Response_201()
	} else {
		fmt.Println(result.Error)
		response.Response_400()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
