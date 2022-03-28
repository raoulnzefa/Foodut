package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	model "github.com/Foodut/backend/modules/product/domain/model"
	srvc "github.com/Foodut/backend/modules/product/domain/service"
	dto "github.com/Foodut/backend/modules/product/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

/**
  Regular "GET" without any extra queries
  Retrieve all registered products from database.
*/
func GetAllProducts(writer http.ResponseWriter, req *http.Request) {

	// Check id query
	productId := req.URL.Query()["id"]

	// Get list of product object
	var products []model.Product = srvc.SearchById(productId)

	// Set response
	var response rspn.Response
	if len(products) > 0 {
		response.Response_200(products)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  Extra "query" using "product name".
  Retrieve all matched product by product name.
  If it is none exact name, then use 'LIKE' sql.
*/
func GetProductByName(writer http.ResponseWriter, req *http.Request) {

	// Check product_name query
	nameToFind := req.URL.Query()["product_name"]

	// Get products using query
	var products []model.Product = srvc.SearchByName(nameToFind)

	// Set response
	var response rspn.Response
	if len(products) > 0 {
		response.Response_200(products)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func GetProductByCategoryId(writer http.ResponseWriter, req *http.Request) {

	// Check category_id query
	nameToFind := req.URL.Query()["category_id"]

	// Get products using query
	var products []model.Product = srvc.SearchByCategoryId(nameToFind)

	// Set response
	var response rspn.Response
	if len(products) > 0 {
		response.Response_200(products)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func GetProductByCategoryName(writer http.ResponseWriter, req *http.Request) {

	// Check category_name query
	nameToFind := req.URL.Query()["category_name"]

	// Get products using query
	var products []model.Product = srvc.SearchByCategoryName(nameToFind)

	// Set response
	var response rspn.Response
	if len(products) > 0 {
		response.Response_200(products)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func DeleteProductById(writer http.ResponseWriter, req *http.Request) {
	// Check id query
	vars := mux.Vars(req)
	productId := vars["id"]

	deleteErr := srvc.DeleteById(productId)
	fmt.Println(deleteErr)
	var response rspn.Response
	//response.Response_200("masuk delete prod ctrl")
	if deleteErr.Error == nil {
		response.Response_200("data has been deleted")
	} else {
		response.Response_400(deleteErr)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  Seller can add a single product each post
  and assign it to his/her store.
*/
func PostProduct(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postProductDto dto.PostProduct
	err := json.NewDecoder(req.Body).Decode(&postProductDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.SendForCreateProduct(postProductDto)

	// Set response
	var response rspn.Response
	if result.Error == nil {
		response.Response_201()
	} else {
		fmt.Println(result.Error)
		response.Response_400("")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
