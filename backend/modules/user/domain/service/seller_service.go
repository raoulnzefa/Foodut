package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
)

func CombinedSellerProduct(storeName []string) model.Seller {
	return repo.FindSellerByStoreName(storeName)
}

// func EmptySearchBy() []model.Seller {
// 	return repo.FindAllProducts()
// }

// func SearchByName(name []string) []model.Product {
// 	products := repo.FindProductsByName(name)

// 	// If none, try find using LIKE
// 	if len(products) < 1 {
// 		products = repo.FindProductsByNameAlike(name)
// 	}

// 	return products
// }
