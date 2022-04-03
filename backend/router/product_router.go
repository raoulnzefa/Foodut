package router

import (
	prController "github.com/Foodut/backend/modules/product/rest-api/controller"
	usrController "github.com/Foodut/backend/modules/user/rest-api/controller"
	"github.com/gorilla/mux"
)

func ProductRouter(router *mux.Router) {
	// CATEGORIES
	//
	//- Insert Categories
	router.HandleFunc("/categories", prController.PostCategory).Methods("POST")

	//- Update Categories
	router.HandleFunc("/categories/{prodcat}", prController.UpdateCategories).Methods("PUT")

	//  PRODUCT
	//
	//- Get All Product
	router.HandleFunc("/products", prController.GetAllProducts).Methods("GET")

	//- Get Products By Name
	router.HandleFunc("/products-by-name", prController.GetProductByName).Methods("GET")

	//- Get Products By Category Id
	router.HandleFunc("/products-by-category-id", prController.GetProductByCategoryId).Methods("GET")

	//- Get Products By Category Name
	router.HandleFunc("/products-by-category-name", prController.GetProductByCategoryName).Methods("GET")

	//- Insert Product
	router.HandleFunc("/products", usrController.Authenticate(prController.PostProduct, 2)).Methods("POST")

	//- Update Product
	router.HandleFunc("/products/{product_id}", prController.EditProduct).Methods("PUT")

	//- Delete Product
	router.HandleFunc("/products/{id}", usrController.Authenticate(prController.DeleteProductById, 2)).Methods("DELETE")

}
