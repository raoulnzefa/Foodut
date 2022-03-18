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

	//- Get Products By Name
	router.HandleFunc("/products-by-name", prController.GetProductByName).Methods("GET")

	//- Insert Product
	router.HandleFunc("/products", prController.PostProduct).Methods("POST")

	//- Update Product
	// router.HandleFunc("/products/{product_id}", prController.UpdateProduct).Methods("PUT")

	//- Delete Product
	// router.HandleFunc("/products/{product_id}", prController.DeleteProduct).Methods("DELETE")

}
