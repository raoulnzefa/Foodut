package services

import (
	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
	dto "github.com/Foodut/backend/modules/product/rest-api/dto"

	"gorm.io/gorm"
)

func SendForCreateCategory(cat dto.Category) *gorm.DB {
	return repo.CreateCategory(MapToCategories(cat))
}

func MapToCategories(cat dto.Category) model.Category {
	// Parse from JSON DTO -> Database Model
	category := model.Category{
		ProductCategory: cat.ProductCategory,
	}

	return category
}
func SendCategoriesForUpdate(cat dto.Category, lastName string) *gorm.DB {
	return repo.UpdateCategories(MapToCategories(cat), lastName)
}
