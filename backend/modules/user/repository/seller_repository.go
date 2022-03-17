package repository

import (
	"fmt"

	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
)

func FindAllSeller() []model.Seller {
	// Check connection
	con := dbController.GetConnection()

	// Get sellers from database, filter by store_name
	var sellers []model.Seller

	con.Find(&sellers)

	return sellers
}

func FindSellerByStoreName(storeName []string) model.Seller {
	// Check connection
	con := dbController.GetConnection()
	var result *gorm.DB

	// Get seller from database, filter by store_name
	var seller model.Seller
	if storeName != nil {
		result = con.Where("store_name = ?", storeName[0]).First(&seller)
	}

	// Association
	if result.Error == nil {
		// with user data
		con.Model(&seller).Association("User").Find(&seller.User)

		// with list product
		con.Model(&seller).Association("ListProduct").Find(&seller.ListProduct)

		// for each product
		for i := 0; i < len(seller.ListProduct); i++ {
			// with picture
			con.Model(&seller.ListProduct).Association("Picture").Find(&seller.ListProduct[i].Picture)
		}
	} else {
		fmt.Println(result.Error)
	}

	return seller
}
