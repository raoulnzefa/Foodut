package services

import (
	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
	"gorm.io/gorm"
)

func SearchById(productId []string) []model.Product {
	products := repo.FindAllProducts(productId)

	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}

func SearchByName(name []string) []model.Product {
	products := repo.FindProductsByName(name)

	// If none, try find using 'LIKE'
	if len(products) < 1 {
		products = repo.FindProductsByNameAlike(name)
	}

	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}

func DeleteById(productId string) *gorm.DB {
	deleteFeedback := repo.DeleteProductById(productId)

	return deleteFeedback
}
