package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	model "github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	"github.com/Foodut/backend/modules/user/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Decode JSON
	var loginUserDto dto.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUserDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cekEmail := GetEmailFromToken(r)
	var user model.User
	var response rspn.Response
	if cekEmail == "" {
		result := srvc.CheckUserLogin(loginUserDto.Email, loginUserDto.Password)
		err := result.Error
		if err == nil {
			generateToken(w, loginUserDto.Email, user.Level)

			response.Response_200("Success Login")
		} else {
			response.Response_404()
		}
	} else {
		res := "Already logged in as " + cekEmail
		response.Response_406(res)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	var response rspn.Response
	response.Response_200("Success Logout |  Bye - Bye")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllUsers(writer http.ResponseWriter, req *http.Request) {

	// Check user_id query
	userId := req.URL.Query()["user_id"]

	// Get list of user object
	var users []model.User = srvc.SearchUserById(userId)

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

// TODO?
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
