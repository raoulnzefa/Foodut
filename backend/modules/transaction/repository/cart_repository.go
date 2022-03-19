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

func UpdateCarts(carts []model.Cart) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Update object in database
	var result *gorm.DB

	// For each data in carts, update its quantity
	for i := 0; i < len(carts); i++ {
		result = con.Model(&model.Cart{}).
			Where("customer_user_id = ? AND product_id = ?", carts[i].CustomerUserID, carts[i].ProductID).
			Update("quantity", carts[i].Quantity)

		if result.Error != nil {
			return result
		}
	}

	return result
}

func DeleteCartById(customerId int) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	var cart []model.Cart
	result := con.Where("customer_user_id = ?", customerId).Delete(&cart)
	return result
}
