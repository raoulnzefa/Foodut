package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/Foodut/backend/modules/transaction/domain/model"
	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
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

func DeleteTransaction(writer http.ResponseWriter, req *http.Request) {
	// Check id query
	vars := mux.Vars(req)
	transId := vars["id"]

	deleteErr := srvc.DeleteById(transId)
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
