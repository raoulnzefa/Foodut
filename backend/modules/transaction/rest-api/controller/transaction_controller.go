package controller

import (
	"encoding/json"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	"github.com/Foodut/backend/modules/transaction/domain/model"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	db := dbController.GetConnection()
	//defer db.Close()

	var transactions []model.Transaction

	transactionID := r.URL.Query()["id"]
	if transactionID != nil {
		db.Where("transaction_id = ?", transactionID[0]).First(&transactions)
	} else {
		db.Find(&transactions)
	}

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
