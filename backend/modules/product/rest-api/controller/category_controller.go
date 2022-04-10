package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	srvc "github.com/Foodut/backend/modules/product/domain/service"
	dto "github.com/Foodut/backend/modules/product/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func PostCategory(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postCategoryDto dto.Category
	err := json.NewDecoder(req.Body).Decode(&postCategoryDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendForCreateCategory(postCategoryDto)

	// Set response
	var response rspn.Response
	if result.Error == nil {
		response.Response_201("Success post category")
	} else {
		response.Response_400("Post category failed")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func UpdateCategories(writer http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	lastName := vars["prodcat"]

	// Decode JSON
	var updateCat dto.Category
	err := json.NewDecoder(req.Body).Decode(&updateCat)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendCategoriesForUpdate(updateCat, lastName)

	// Set response
	var response rspn.Response
	if result.Error == nil {
		response.Response_201("Success update category")
	} else {
		response.Response_400(result.Error)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
