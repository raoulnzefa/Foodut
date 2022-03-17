package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/product/domain/model"
)

func GetProductsAssociation(products []model.Product) {
	// Check Connection
	con := dbController.GetConnection()

	// For each product
	for i := 0; i < len(products); i++ {
		// with picture
		for j := 0; j < len(products); j++ {
			con.Model(&products).Association("Picture").Find(&products[i].Picture)
		}
	}

}

func FindAllProducts(productId []string) []model.Product {
	// Check Connection
	con := dbController.GetConnection()

	// Get all products from database
	var products []model.Product

	// Extra query by Id
	if productId != nil {
		con.Find(&products, productId[0])
	} else {
		con.Find(&products)
	}

	return products
}

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
