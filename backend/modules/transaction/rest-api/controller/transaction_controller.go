package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/transaction/domain/model"
	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
	rspn "github.com/Foodut/backend/responses"
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
