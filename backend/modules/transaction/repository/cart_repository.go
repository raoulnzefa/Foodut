package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	"gorm.io/gorm"
)

func CreateCart(carts []model.Cart) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&carts)

	return result
}

func GetCartById(customerId int) []model.Cart {
	// Check connection
	con := dbController.GetConnection()

	var cart []model.Cart
	con.Where("customer_user_id = ?", customerId).Find(&cart)
	return cart
}

func DeleteCartById(customerId int) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	var cart []model.Cart
	result := con.Where("customer_user_id = ?", customerId).Delete(&cart)
	return result
}
