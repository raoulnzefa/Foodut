package repository

import (
	"fmt"

	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/product/domain/model"
	"gorm.io/gorm"
)

func FindAllProducts(productId []string) []model.Product {
	// Check Connection
	con := dbController.GetConnection()
	var result *gorm.DB

	// Get all products from database
	var products []model.Product

	// Extra query by Id
	if productId != nil {
		result = con.Find(&products, productId[0])
	} else {
		result = con.Find(&products)
	}

	// Association
	if result.Error == nil && len(products) > 0 {
		// for each product
		for i := 0; i < len(products); i++ {
			// with picture
			for j := 0; j < len(products); j++ {
				con.Model(&products).Association("Picture").Find(&products[i].Picture)
			}
		}
	} else {
		fmt.Println(result.Error)
	}

	return products
}

func FindProductsByName(productName []string) []model.Product {
	// Check Connection
	con := dbController.GetConnection()
	var result *gorm.DB

	// Get product from database, filter by product_name
	var products []model.Product
	if productName != nil {
		result = con.Where("product_name = ?", productName[0]).Find(&products)
	}

	// Association
	if result.Error == nil && len(products) > 0 {
		// for each product
		for i := 0; i < len(products); i++ {
			// with picture
			for j := 0; j < len(products[i].Picture); j++ {
				con.Model(&products[i].Picture).Association("Picture").Find(&products[i].Picture)
			}
		}
	} else {
		fmt.Println(result.Error)
	}

	return products
}

func FindProductsByNameAlike(productName []string) []model.Product {
	// Check Connection
	con := dbController.GetConnection()
	var result *gorm.DB

	// Get product from database, filter by name LIKE
	var products []model.Product
	if productName != nil {
		like := "%" + productName[0] + "%"
		result = con.Where("product_name LIKE ?", like).Find(&products)
	}

	// Association
	if result.Error == nil && len(products) > 0 {
		// for each product
		for i := 0; i < len(products); i++ {
			// with picture
			for j := 0; j < len(products[i].Picture); j++ {
				con.Model(&products[i].Picture).Association("Picture").Find(&products[i].Picture)
			}
		}
	} else {
		fmt.Println(result.Error)
	}

	return products
}
