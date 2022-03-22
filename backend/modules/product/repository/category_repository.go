package repository

import (
	dbController "github.com/Foodut/backend/database"
	"github.com/Foodut/backend/modules/product/domain/model"
)

func GetCategoryByName(categoryName []string) model.Category {
	// Check connection
	con := dbController.GetConnection()

	var category model.Category
	// Get category from database, check by its name
	if categoryName != nil {
		con.Where("product_category = ?", categoryName[0]).Find(&category)
	}

	return category

}
