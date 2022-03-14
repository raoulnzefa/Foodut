package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/product/domain/model"
)

func FindProductByName(productName []string) []model.Product {
	// Check Connection
	con := dbController.GetConnection()

	// Get product from database, if there is some query, do filter
	var products []model.Product
	if productName != nil {
		con.Where("product_name = ?", productName[0]).First(&products)
	} else {
		con.Find(&products)
	}

	return products
}
