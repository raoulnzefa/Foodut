package services

import (
	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
)

func SearchById(productId []string) []model.Product {
	return repo.FindAllProducts(productId)
}

func SearchByName(name []string) []model.Product {
	products := repo.FindProductsByName(name)

	// If none, try find using LIKE
	if len(products) < 1 {
		products = repo.FindProductsByNameAlike(name)
	}

	return products
}
