package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
)

func EmptySellerMinimal() []dto.SellerMinimal {
	sellers := repo.FindAllSeller()

	var sellerDto []dto.SellerMinimal

	// Move from database model to JSON DTO
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
	seller := repo.FindSellerByStoreName(storeName)

	// Association
	if seller.UserID != 0 {
		repo.GetOneSellerAssociation(&seller)
	}

	return seller
}
