package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
)

func EmptySellerMinimal() []dto.SellerMinimal {
	sellers := repo.FindAllSeller()

	var sellerDto []dto.SellerMinimal

	for _, s := range sellers {
		var tempSeller dto.SellerMinimal
		tempSeller.SetUserId(s.UserID)
		tempSeller.SetStoreName(s.StoreName)
		tempSeller.SetCity(s.City)
		sellerDto = append(sellerDto, tempSeller)
	}

	return sellerDto
}

func SearchByStoreName(storeName []string) model.Seller {
	return repo.FindSellerByStoreName(storeName)
}
