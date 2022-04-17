package repository

import (
	dbController "github.com/Foodut/backend/database"
	"github.com/Foodut/backend/modules/product/domain/model"
	"gorm.io/gorm"
)

func ReadOneProductPicture(productId int) []string {
	// Check connection
	con := dbController.GetConnection()

	var paths []string

	if productId > 0 {
		con.Raw(
			"SELECT picture_path AS product_picture "+
				"FROM `pictures` "+
				"WHERE product_id = ?", productId).
			Scan(&paths)
	}

	return paths
}

func DeleteProductPictures(productId string) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	result := con.Where("product_id = ?", productId).Delete(&model.Picture{})

	return result
}
