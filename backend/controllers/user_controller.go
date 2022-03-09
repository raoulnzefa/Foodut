package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Foodut/backend/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	// dbSQL, ok := db.DB()
	// if ok != nil {
	// 	defer dbSQL.Close()
	// }

	var users []models.User

	name := r.URL.Query()["name"]
	if name != nil {
		db.Where("name = ?", name[0]).First(&users)
	} else {
		db.Find(&users)
	}

	// Set response
	var response models.UserResponse
	if len(users) > 0 {
		response.Status = 200
		response.Message = "Success Get User Data"
		response.Data = users
	} else {
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
