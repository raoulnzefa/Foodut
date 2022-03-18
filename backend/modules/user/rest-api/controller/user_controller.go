package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	// Get list of user object
	var users []model.User = srvc.EmptyUserSearch()

	// Set response
	var response rspn.Response
	if len(users) > 0 {
		response.Response_200(users)
	} else {
		response.Response_204()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// TODO
// func GetUsersByUsernameLike(w http.ResponseWriter, r *http.Request) {

// 	// Get list of user object
// 	var users []model.User = srvc.EmptySearchBy()

// 	name := r.URL.Query()["name"]
// 	if name != nil {
// 		db.Where("name = ?", name[0]).First(&users)
// 	} else {
// 		db.Find(&users)
// 	}

// 	// Set response
// 	var response rspn.Response
// 	if len(users) > 0 {
// 		response.Response_200(users)
// 	} else {
// 		response.Response_204()
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func DeleteUser(writer http.ResponseWriter, req *http.Request) {
	// Check id query
	vars := mux.Vars(req)
	userId := vars["id"]

	deleteErr := srvc.DeleteById(userId)
	fmt.Println(deleteErr)
	var response rspn.Response
	//response.Response_200("masuk delete prod ctrl")
	if deleteErr.Error == nil {
		response.Response_200("data has been deleted")
	} else {
		response.Response_400(deleteErr)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
