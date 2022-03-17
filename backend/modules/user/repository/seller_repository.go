package repository

import (
	dbController "github.com/Foodut/backend/database"
	prRepo "github.com/Foodut/backend/modules/product/repository"
	model "github.com/Foodut/backend/modules/user/domain/model"
)

func GetSellersAssociation(sellers []model.Seller) {
	// For each seller
	for i := 0; i < len(sellers); i++ {
		GetOneSellerAssociation(&sellers[i])
	}
}

func GetOneSellerAssociation(seller *model.Seller) {
	// Check connection
	con := dbController.GetConnection()

	// with user data
	con.Model(&seller).Association("User").Find(&seller.User)

	// with list product
	con.Model(&seller).Association("ListProduct").Find(&seller.ListProduct)

	// get each product association
	prRepo.GetProductsAssociation(seller.ListProduct)
}

func FindAllSeller() []model.Seller {
	// Check connection
	con := dbController.GetConnection()

	// Get sellers from database
	var sellers []model.Seller

	con.Find(&sellers)

	return sellers
}

func FindSellerByStoreName(storeName []string) model.Seller {
	// Check connection
	con := dbController.GetConnection()

	// Get seller from database, filter by store_name
	var seller model.Seller
	if storeName != nil {
		con.Where("store_name = ?", storeName[0]).First(&seller)
	}

	return seller
}
