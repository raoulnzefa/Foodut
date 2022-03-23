package router

import (
	trController "github.com/Foodut/backend/modules/transaction/rest-api/controller"
	usrController "github.com/Foodut/backend/modules/user/rest-api/controller"
	"github.com/gorilla/mux"
)

func TransactionRouter(router *mux.Router) {

	//  TRANSACTION
	//
	//- Get All Transaction
	router.HandleFunc("/transactions", usrController.Authenticate(trController.GetAllTransactions, 3)).Methods("GET")

	//- Get Transaction Detail
	// router.HandleFunc("/transactions-detail", trController.GetTransactionDetail).Methods("GET")

	//- Get Carts With Availability
	router.HandleFunc("/cart-a", trController.GetCartWithAvailability).Methods("GET")

	//- Insert Cart
	router.HandleFunc("/cart", usrController.Authenticate(trController.PostToCart, 1)).Methods("POST")

	//- Insert Transaction
	router.HandleFunc("/transactions", usrController.Authenticate(trController.PostTransaction, 1)).Methods("POST")

	//- Update Cart
	router.HandleFunc("/cart", usrController.Authenticate(trController.UpdateCart, 1)).Methods("PUT")

	//- Delete Carts
	router.HandleFunc("/cart", trController.DeleteCarts).Methods("DELETE")

	//- Delete Specific Product from a Cart (Update w/ delete spesific item)
	router.HandleFunc("/cart-specific", trController.DeleteSpesificProductFromCart).Methods("DELETE")

	//- Update Transaction
	// router.HandleFunc("/transactions/{transaction_id}", trController.UpdateTransaction).Methods("PUT")

	//- Delete Transaction
	router.HandleFunc("/transactions/{id}", usrController.Authenticate(trController.DeleteTransaction, 3)).Methods("DELETE")
}
