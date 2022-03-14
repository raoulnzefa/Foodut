package router

import (
	trController "github.com/Foodut/backend/modules/transaction/rest-api/controller"
	"github.com/gorilla/mux"
)

func TransactionRouter(router *mux.Router) {

	//  TRANSACTION
	//
	//- Get All Transaction
	router.HandleFunc("/transactions", trController.GetAllTransactions).Methods("GET")
}
