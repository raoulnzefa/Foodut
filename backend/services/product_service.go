package services

import (
	model "github.com/Foodut/backend/models"
	repo "github.com/Foodut/backend/repositories"
)

func SearchByName(name []string) []model.Product {
	return repo.FindProductByName(name)
}
