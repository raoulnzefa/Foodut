package router

import (
	usrController "github.com/Foodut/backend/modules/user/rest-api/controller"
	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router) {
	// router.HandleFunc("/usersTest", usrController.GetUserResponse).Methods("GET")
	// router.HandleFunc("/TestDB", usrController.TestDB).Methods("GET")

	//  USER
	//
	//- Get All User
	router.HandleFunc("/users", usrController.GetAllUsers).Methods("GET")

	//- Get All Seller
	router.HandleFunc("/sellers", usrController.GetAllSeller).Methods("GET")

	//- Get Seller with Products
	router.HandleFunc("/store", usrController.GetSellerByStoreWithProducts).Methods("GET")

	//- Get User Detail Transaction
	// router.HandleFunc("/detail-user-transactions", usrController.GetDetailUserTransaction).Methods("GET")

	//- Post User
	// router.HandleFunc("/user", usrController.PostUser).Methods("POST")

	//- Post Customer
	router.HandleFunc("/customer", usrController.PostCustomer).Methods("POST")

	//- Post Seller
	router.HandleFunc("/seller", usrController.PostSeller).Methods("POST")

	//- Post Admin
	router.HandleFunc("/admin", usrController.PostAdmin).Methods("POST")

	//- Update User
	// router.HandleFunc("/users/{user_id}", usrController.UpdateUser).Methods("PUT")

	//- Delete User
	router.HandleFunc("/users/{id}", usrController.DeleteUser).Methods("DELETE")

	//- Login User
	router.HandleFunc("/loginUser", usrController.LoginUser).Methods("GET")

	//- Logout User
	router.HandleFunc("/logout", usrController.Logout).Methods("GET")

}
