package repository

import dbController "github.com/Foodut/backend/database"

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
