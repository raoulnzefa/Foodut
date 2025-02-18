package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
)

func ReadAllUsers(userId []string) []model.User {
	// Check connection
	con := dbController.GetConnection()

	// Get users from database, filter by store_name
	var users []model.User

	// Extra query by id
	if userId != nil {
		con.Find(&users, userId[0])
	} else {
		con.Find(&users)
	}

	return users
}

func CreateUser(user model.User) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&user)

	return result
}

func DeleteUserById(userId string) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	//delete user on db by id
	var user model.User
	result := con.Delete(&user, userId)
	return result
}

func CheckUserEmailPassword(email string, password string) (model.User, *gorm.DB) {
	con := dbController.GetConnection()
	var user model.User
	result := con.Where("email = ? AND password = ?", email, password).First(&user)
	return user, result
}

func SaveUser(user model.User) *gorm.DB {
	con := dbController.GetConnection()
	result := con.Save(&user)
	return result
}
