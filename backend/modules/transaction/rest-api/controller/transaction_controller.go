package controller

import (
	"encoding/json"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
	dto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
	"github.com/gorilla/mux"
)

func GetAllTransactions(writer http.ResponseWriter, req *http.Request) {

	transactionId := req.URL.Query()["id"]

	var transactions []dto.Transaction = srvc.SearchById(transactionId)

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

func GetAllOrders(writer http.ResponseWriter, req *http.Request) {

	sellerId := req.URL.Query()["sellerId"]

	// var transactions []dto.Transaction = srvc.SearchById(transactionId)

	con := dbController.GetConnection()

	var OrderDetail []dto.OrderDetail

	con.Raw(
		"SELECT transaction_details.*, products.product_name, transactions.customer_id, "+
			"transactions.payment_option, transactions.transaction_date, "+
			"users.name, transactions.address "+
			"FROM `transactions`, `transaction_details` , `products`, `users` "+
			"WHERE users.id = transactions.customer_id "+
			"AND transactions.id = transaction_details.transaction_id "+
			"AND transaction_details.product_id = products.id "+
			"AND status = 'ORDER' "+
			"AND transaction_details.seller_id = ?",
		sellerId[0]).
		Scan(&OrderDetail)

	// Set response
	var response rspn.Response
	if len(OrderDetail) > 0 {
		response.Response_200(OrderDetail)

	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

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
		response.Response_201()
	} else {
		response.Response_400(isError.Error())
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
