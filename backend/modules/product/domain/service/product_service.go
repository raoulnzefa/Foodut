package services

import (
	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
)

func SearchById(productId []string) []model.Product {
	// Get products from database table
	products := repo.FindAllProducts(productId)

	// Fill out relation
	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}

func SearchByName(name []string) []model.Product {
	products := repo.FindProductsByName(name)

	// If none, try find using LIKE
	if len(products) < 1 {
		products = repo.FindProductsByNameAlike(name)
	}

	// Fill out relation
	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}
