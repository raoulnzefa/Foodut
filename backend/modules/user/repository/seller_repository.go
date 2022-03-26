package repository

import (
	dbController "github.com/Foodut/backend/database"
	prRepo "github.com/Foodut/backend/modules/product/repository"
	model "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
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

func ReadAllSeller() []model.Seller {
	// Check connection
	con := dbController.GetConnection()

	// Get sellers from database
	var sellers []model.Seller

	con.Find(&sellers)

	return sellers
}

func ReadSellerByStoreName(storeName []string) model.Seller {
	// Check connection
	con := dbController.GetConnection()

	// Get seller from database, filter by store_name
	var seller model.Seller
	if storeName != nil {
		con.Where("store_name = ?", storeName[0]).First(&seller)
	}

	return seller
}

func CreateSeller(seller model.Seller) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&seller)

	return result
}

func DeleteSellerByCustId(userId string) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	//delete Customer on db by id
	var sell model.Seller
	result := con.Where("user_id = ?", userId).Delete(&sell)
	return result
}
