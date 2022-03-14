package router

import (
	prController "github.com/Foodut/backend/modules/product/rest-api/controller"
	"github.com/gorilla/mux"
)

func ProductRouter(router *mux.Router) {

	//  PRODUCT
	//
	//- Get All Product
	router.HandleFunc("/products", prController.GetAllProducts).Methods("GET")

}
