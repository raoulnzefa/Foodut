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

func GetCartByCustId(customerId int) []model.Cart {
	// Check connection
	con := dbController.GetConnection()

	var cart []model.Cart
	con.Where("customer_user_id = ?", customerId).Find(&cart)
	return cart
}

func GetCartByCustIdWithAvailablityCheck(customerId int) []model.Cart {
	// Check connection
	con := dbController.GetConnection()

	var carts []model.Cart
	con.Raw(
		"SELECT carts.customer_user_id, carts.product_id, carts.quantity "+
			"FROM carts, products "+
			"WHERE carts.product_id = products.id "+
			"AND carts.quantity <= products.product_stock "+
			"AND carts.customer_user_id = ?",
		customerId).
		Scan(&carts)
	return carts
}

func UpdateCarts(carts []model.Cart) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	result := con.Save(&carts)

	return result
}

func DeleteCarts(carts []model.Cart) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	result := con.Delete(&carts)

	return result
}

func DeleteCartByCustId(customerId int) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	var cart []model.Cart
	result := con.Where("customer_user_id = ?", customerId).Delete(&cart)
	return result
}

func DeleteCartByCustIdAndProductId(customerId int, productId int) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	var cart []model.Cart
	result := con.Where(
		"customer_user_id = ? AND product_id = ?", customerId, productId).
		Delete(&cart)
	return result
}
