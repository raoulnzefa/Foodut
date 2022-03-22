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

	//- Get Transaction Detail
	// router.HandleFunc("/transactions-detail", trController.GetTransactionDetail).Methods("GET")

	//- Insert Cart
	router.HandleFunc("/cart", trController.PostToCart).Methods("POST")

	//- Insert Transaction
	router.HandleFunc("/transactions", trController.PostTransaction).Methods("POST")

	//- Update Cart
	router.HandleFunc("/cart", trController.UpdateCart).Methods("PUT")

	//- Delete Specific Cart (Update w/ delete some item)
	router.HandleFunc("/cart-specific", trController.DeleteSpesificCart).Methods("DELETE")

	//- Update Transaction
	// router.HandleFunc("/transactions/{transaction_id}", trController.UpdateTransaction).Methods("PUT")

	//- Delete Transaction
	router.HandleFunc("/transactions/{id}", trController.DeleteTransaction).Methods("DELETE")
}
