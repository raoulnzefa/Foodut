package router

import (
	usrController "github.com/Foodut/backend/modules/user/rest-api/controller"
	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router) {
	router.HandleFunc("/usersTest", usrController.GetUserResponse).Methods("GET")

	//  USER
	//
	//- Get All User
	router.HandleFunc("/users", usrController.GetAllUsers).Methods("GET")

}
