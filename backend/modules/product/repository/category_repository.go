package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/product/domain/model"
	"gorm.io/gorm"
)

func ReadAllCategory() []model.Category {
	// Check connection
	con := dbController.GetConnection()

	var category []model.Category

	con.Find(&category)

	return category
}

func ReadCategoryByName(categoryName []string) model.Category {
	// Check connection
	con := dbController.GetConnection()

	var category model.Category
	// Get category from database, check by its name
	if categoryName != nil {
		con.Where("product_category = ?", categoryName[0]).Find(&category)
	}

	return category

}

func CreateCategory(cat model.Category) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&cat)

	return result
}

func UpdateCategories(cat model.Category, lastName string) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()
	//var cats model.Category
	result := con.Model(&cat).Where("product_category = ?", lastName).Update("product_category", cat.ProductCategory)

	return result
}
