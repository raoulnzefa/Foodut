package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
)

func FindAllUsers() []model.User {
	// Check connection
	con := dbController.GetConnection()

	// Get users from database, filter by store_name
	var users []model.User

	con.Find(&users)

	return users
}

func CreateUser(user model.User) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&user)

	return result
}

func CreateCustomer(customer model.Customer) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&customer)

	return result
}

func CreateSeller(seller model.Seller) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&seller)

	return result
}

func CreateAdmin(admin model.Admin) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&admin)

	return result
}
