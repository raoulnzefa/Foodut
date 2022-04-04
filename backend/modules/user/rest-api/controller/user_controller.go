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

/**
  User can login using email and password
*/
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Decode JSON
	var loginUserDto dto.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUserDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cekEmail := GetEmailFromToken(r)
	var response rspn.Response
	if cekEmail == "" {
		user, result := srvc.CheckUserLogin(loginUserDto.Email, loginUserDto.Password)
		err := result.Error
		if err == nil {
			generateToken(w, user.ID, loginUserDto.Email, user.Level)
			response.Response_200(user.ID)
		} else {
			response.Response_404("Login fail || " + err.Error())
		}
	} else {
		res := "Already logged in as " + cekEmail
		response.Response_406(res)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/**
  Logging out user by reset
  its token
*/
func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	var response rspn.Response
	response.Response_200("Success Logout |  Bye - Bye")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/**
  Admin can use this method to get all user
  or add user_id as query param to filter it
  by user_id
*/
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
		response.Response_204("Get user fail")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  User can use this method to get its profile
*/
func GetUserById(writer http.ResponseWriter, req *http.Request) {
	// Check id query
	vars := mux.Vars(req)
	userIds := [1]string{vars["user_id"]}

	// Get list of user object
	var users []model.User = srvc.SearchUserById(userIds[:])

	// Set response
	var response rspn.Response
	if len(users) > 0 {
		response.Response_200(users)
	} else {
		response.Response_204("Get user by id fail")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  Admin can use this method to delete a user
  Search user by id and delete it
*/
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

/**
  User can edit its profile anytime
*/
func EditUser(writer http.ResponseWriter, req *http.Request) {
	// Decode JSON
	var editUserDto dto.EditUser
	err := json.NewDecoder(req.Body).Decode(&editUserDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Get id from path
	vars := mux.Vars(req)
	userID := vars["user_id"]

	var id []string
	id = append(id, userID)
	users := srvc.SearchUserById(id)

	result1, result2, level := srvc.EditUser(users[0], editUserDto)
	var response rspn.Response
	// Custsomer = 1 dan Seller = 2
	if level == 1 {
		if err == nil {
			if result1.Error == nil && result2.Error == nil {
				response.Response_200("Edit Customer Data Success")
			} else {
				response.Response_400("Edit Customer Data Failed " + result1.Error.Error() + " " + result2.Error.Error())
			}
		} else {
			response.Response_400("Edit Customer Data Failed, ID Not Valid" + err.Error())
		}
	} else if level == 2 {

		if err == nil {
			if result1.Error == nil && result2.Error == nil {
				response.Response_200("Edit Seller Data Success")
			} else {
				response.Response_400("Edit Seller Data Failed " + result1.Error.Error() + " " + result2.Error.Error())
			}
		} else {
			response.Response_400("Edit Seller Data Failed, ID Not Valid" + err.Error())
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
