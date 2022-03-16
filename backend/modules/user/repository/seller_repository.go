package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
)

func FindSellerByStoreName(storeName []string) model.Seller {
	// Check connection
	con := dbController.GetConnection()

	// Get seller from database, filter by store_name
	var seller model.Seller
	if storeName != nil {
		con.Where("store_name = ?", storeName).First(&seller)
	}

	// Get user data
	con.Model(&seller).Association("User")
	con.Model(&seller).Association("User").Find(&seller.User)

	// Get seller's products
	con.Model(&seller).Association("ListProduct")
	con.Model(&seller).Association("ListProduct").Find(&seller.ListProduct)

	// Get product pictures
	con.Model(&seller.ListProduct).Association("Picture")
	for i := 0; i < len(seller.ListProduct); i++ {
		con.Model(&seller.ListProduct[i]).Association("Picture").Find(&seller.ListProduct[i].Picture)
	}

	return seller
}
