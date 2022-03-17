package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/transaction/domain/model"
	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {

	transactionId := r.URL.Query()["id"]

	var transactions []model.Transaction = srvc.SearchById(transactionId)

	// Set response
	var response rspn.Response
	if len(transactions) > 0 {
		response.Response_200(transactions)

	} else {
		response.Response_204()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
