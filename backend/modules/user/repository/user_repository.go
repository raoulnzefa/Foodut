package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
)

func FindAllUsers() []model.User {
	// Check connection
	con := dbController.GetConnection()

	// Get users from database, filter by store_name
	var users []model.User

	con.Find(&users)

	return users
}
