package controller

import (
	"encoding/json"
	"net/http"

	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
	dto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
	"github.com/gorilla/mux"
)

/**
  Admin can get all transaction of customer
  or can filter it by customer id
  to get a spesific transction
*/
func GetAllTransactions(writer http.ResponseWriter, req *http.Request) {

	transactionId := req.URL.Query()["id"]

	var transactions []dto.Transaction = srvc.SearchAllOrById(transactionId)

	// Set response
	var response rspn.Response
	if len(transactions) > 0 {
		response.Response_200(transactions)
	} else {
		response.Response_204("Get transaction fail")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  Admin can get all transaction of customer
  or can filter it by customer id
  to get a spesific transction
  'DECENTRALIZED JSON'
*/
func GetAllTransactionsDecentralize(writer http.ResponseWriter, req *http.Request) {

	var orders []dto.TransactionDecentralize = srvc.SendForGetTransactionDetail()

	// Set response
	var response rspn.Response
	if len(orders) > 0 {
		response.Response_200(orders)

	} else {
		response.Response_204("Get transaction fail")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  Customer can get his/her transaction
*/
func GetCustomerTransactions(writer http.ResponseWriter, req *http.Request) {

	customerId := req.URL.Query()["customerId"]

	var transactions []dto.Transaction = srvc.SearchByCustId(customerId)

	// Set response
	var response rspn.Response
	if len(transactions) > 0 {
		response.Response_200(transactions)
	} else {
		response.Response_204("Get transaction fail")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  Customer can get his/her all transaction
  'DECENTRALIZED JSON'
*/
func GetCustomerTransactionsDecentralize(writer http.ResponseWriter, req *http.Request) {

	customerId := req.URL.Query()["customerId"]

	var orders []dto.TransactionDecentralize = srvc.SendCustomerIdForGetTransactionDetail(customerId)

	// Set response
	var response rspn.Response
	if len(orders) > 0 {
		response.Response_200(orders)

	} else {
		response.Response_204("Get transaction fail")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  Seller can view all ordered
  product related to his/her.
  Filtered by 'ORDER' status
  in transaction_details schema
*/
func GetAllOrders(writer http.ResponseWriter, req *http.Request) {

	sellerId := req.URL.Query()["sellerId"]

	var orders []dto.OrderDetail = srvc.SendSellerIdForReadOrder(sellerId)

	// Set response
	var response rspn.Response
	if len(orders) > 0 {
		response.Response_200(orders)
	} else {
		response.Response_204("Get seller orders fail")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  This method purpose is for customer payment
*/
func PostTransaction(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postTransactionDto dto.PostTransaction
	err := json.NewDecoder(req.Body).Decode(&postTransactionDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	isError := srvc.InsertTransactionAvailabilityCheck(postTransactionDto)

	// Set response
	var response rspn.Response
	if isError == nil {
		response.Response_201("Success post transaction")
	} else {
		response.Response_400(isError.Error())
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  'Rarely happen in any enterprise/company'
  Delete transaction by its id
*/
func DeleteTransaction(writer http.ResponseWriter, req *http.Request) {
	// Check id query
	vars := mux.Vars(req)
	transId := vars["id"]

	deleteErr := srvc.DeleteById(transId)
	var response rspn.Response

	if deleteErr.Error == nil {
		response.Response_200("Success delete transaction")
	} else {
		response.Response_400(deleteErr)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
