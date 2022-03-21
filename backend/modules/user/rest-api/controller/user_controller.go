package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	database "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	rspn "github.com/Foodut/backend/responses"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()

	email := r.URL.Query()["email"]
	password := r.URL.Query()["password"]
	var user model.User
	cekEmail := GetEmailFromToken(r)

	if cekEmail == "" {
		if err := db.Where("email = ? AND password = ?", email[0], password[0]).First(&user).Error; err == nil {
			generateToken(w, user.Email, user.Level)

			sendSuccessResponse(w)
		} else {
			sendErrorResponse(w)
		}
	} else {
		var response rspn.BasicResponse
		response.StatusCode = 406
		response.Message = "Already logged in as " + cekEmail
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	var response rspn.BasicResponse
	response.StatusCode = 200
	response.Message = "Success Logout |  Bye - Bye"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response rspn.BasicResponse
	response.StatusCode = 200
	response.Message = "Success Login"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response rspn.BasicResponse
	response.StatusCode = 400
	response.Message = "Failed Login"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUnauthorizedResponse(w http.ResponseWriter) {
	var response rspn.BasicResponse
	response.StatusCode = 401
	response.Message = "Unauthorized Access"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllUsers(writer http.ResponseWriter, req *http.Request) {

	// Get list of user object
	var users []model.User = srvc.EmptyUserSearch()

	// Set response
	var response rspn.Response
	if len(users) > 0 {
		response.Response_200(users)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

// TODO
// func GetUsersByUsernameLike(writer http.ResponseWriter, req *http.Request) {

// 	// Get list of user object
// 	var users []model.User = srvc.EmptySearchBy()

// 	name := req.URL.Query()["name"]
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

// 	writer.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(writer).Encode(response)
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
