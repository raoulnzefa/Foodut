package services

import (
	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
)

func SearchByName(name []string) []model.Product {
	return repo.FindProductByName(name)
}
