package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/product/domain/model"
)

func FindProductsByName(productName []string) []model.Product {
	// Check Connection
	con := dbController.GetConnection()

	// Get product from database, filter by product_name
	var products []model.Product
	if productName != nil {
		con.Where("product_name = ?", productName[0]).Find(&products)
	}

	return products
}

func FindProductsByNameAlike(productName []string) []model.Product {
	// Check Connection
	con := dbController.GetConnection()

	// Get product from database, filter by name LIKE
	var products []model.Product
	if productName != nil {
		like := "%" + productName[0] + "%"
		con.Where("product_name LIKE ?", like).Find(&products)
	}

	return products
}

func FindAllProducts() []model.Product {
	// Check Connection
	con := dbController.GetConnection()

	// Get all products from database
	var products []model.Product

	con.Find(&products)

	return products
}
