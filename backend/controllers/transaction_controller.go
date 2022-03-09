package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Foodut/backend/models"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	//defer db.Close()

	var transactions []models.Transaction

	transactionID := r.URL.Query()["id"]
	if transactionID != nil {
		db.Where("transaction_id = ?", transactionID[0]).First(&transactions)
	} else {
		db.Find(&transactions)
	}

	// Set response
	var response models.TransactionResponse
	if len(transactions) > 0 {
		response.Status = 200
		response.Message = "Success Get Transaction"
		response.Data = transactions
	} else {
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
