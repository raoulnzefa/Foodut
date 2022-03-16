package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	rspn "github.com/Foodut/backend/responses"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	// Get list of user object
	var users []model.User = srvc.EmptySearchBy()

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
