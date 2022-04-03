package router

import (
	usrController "github.com/Foodut/backend/modules/user/rest-api/controller"
	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router) {

	//  USER
	//- Get All User
	router.HandleFunc("/users", usrController.Authenticate(usrController.GetAllUsers, 3)).Methods("GET")

	//Get User By Id
	router.HandleFunc("/users/{user_id}", usrController.AuthenticateMinimumLevel(usrController.GetUserById, 1)).Methods("GET")

	//- Get All Seller
	router.HandleFunc("/sellers", usrController.GetAllSeller).Methods("GET")

	//- Get Customer With Transaction
	router.HandleFunc("/customers", usrController.GetAllCustomerWithAssociation).Methods("GET")

	//- Get Customer With Transaction Complete
	router.HandleFunc("/customers-complete", usrController.GetAllCustomerWithAssociationComplete).Methods("GET")

	//- Get Seller with Products
	router.HandleFunc("/store", usrController.GetSellerByIdWithProducts).Methods("GET")

	//- Post User
	// router.HandleFunc("/user", usrController.PostUser).Methods("POST")

	//- Post Customer
	router.HandleFunc("/customer", usrController.PostCustomer).Methods("POST")

	//- Post Seller
	router.HandleFunc("/seller", usrController.PostSeller).Methods("POST")

	//- Post Admin
	router.HandleFunc("/admin", usrController.PostAdmin).Methods("POST")

	//- Update User
	//Customer or Seller
	router.HandleFunc("/users/{user_id}", usrController.EditUser).Methods("PUT")

	//- Delete User
	router.HandleFunc("/users/{id}", usrController.DeleteUser).Methods("DELETE")

	//- Login User
	router.HandleFunc("/loginUser", usrController.LoginUser).Methods("POST")

	//- Logout User
	router.HandleFunc("/logout", usrController.Logout).Methods("GET")

}
