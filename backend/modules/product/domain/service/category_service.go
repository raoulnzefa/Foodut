package services

import (
	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
	dto "github.com/Foodut/backend/modules/product/rest-api/dto"

	"gorm.io/gorm"
)

func SendForReadAllCategory() []dto.GetCategory {
	categories := repo.ReadAllCategory()

	var getCategoriesDto []dto.GetCategory
	for i := 0; i < len(categories); i++ {
		getCategoriesDto = append(getCategoriesDto, MapToGetCategoriesDto(categories[i]))
	}

	return getCategoriesDto
}

func SendForCreateCategory(cat dto.Category) *gorm.DB {
	return repo.CreateCategory(MapToCategories(cat))
}

func SendCategoriesForUpdate(cat dto.Category, lastName string) *gorm.DB {
	return repo.UpdateCategories(MapToCategories(cat), lastName)
}

func MapToCategories(cat dto.Category) model.Category {
	// Parse from JSON DTO -> Database Model
	category := model.Category{
		ProductCategory: cat.ProductCategory,
	}

	return category
}

func MapToGetCategoriesDto(cat model.Category) dto.GetCategory {
	// Parse from Database Model -> JSON DTO
	category := dto.GetCategory{
		ID:              cat.ID,
		ProductCategory: cat.ProductCategory,
	}

	return category
}
