package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/transaction/domain/model"
	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
	dto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
	"github.com/gorilla/mux"
)

func GetAllTransactions(writer http.ResponseWriter, req *http.Request) {

	transactionId := req.URL.Query()["id"]

	var transactions []model.Transaction = srvc.SearchById(transactionId)

	// Set response
	var response rspn.Response
	if len(transactions) > 0 {
		response.Response_200(transactions)

	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func PostTransaction(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postTransactionDto dto.Transaction
	err := json.NewDecoder(req.Body).Decode(&postTransactionDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.InsertTransaction(postTransactionDto)

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

func DeleteTransaction(writer http.ResponseWriter, req *http.Request) {
	// Check id query
	vars := mux.Vars(req)
	transId := vars["id"]

	deleteErr := srvc.DeleteById(transId)
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
