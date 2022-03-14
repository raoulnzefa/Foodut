package controller

import (
	"encoding/json"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	usrModel "github.com/Foodut/backend/modules/user/domain/model"
	response "github.com/Foodut/backend/responses"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := dbController.GetConnection()

	var users []usrModel.User

	name := r.URL.Query()["name"]
	if name != nil {
		db.Where("name = ?", name[0]).First(&users)
	} else {
		db.Find(&users)
	}

	// Set response
	var response response.Response
	if len(users) > 0 {
		response.Response_200(users)
	} else {
		response.Response_204()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
