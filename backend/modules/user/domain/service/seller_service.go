package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func EmptySellerMinimal() []dto.SellerMinimal {
	sellers := repo.ReadAllSeller()

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
	seller := repo.ReadSellerByStoreName(storeName)

	// Association
	if seller.UserID > 0 {
		repo.GetOneSellerAssociation(&seller)
	}

	return seller
}

func SendForCreateSeller(u dto.PostSeller) *gorm.DB {
	return repo.CreateSeller(MapToSeller(u))
}

func MapToSeller(u dto.PostSeller) model.Seller {

	// Parse from JSON DTO -> Database Model
	usr := dto.PostUser{
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}

	user := MapToUser(usr, 2)

	sell := model.Seller{
		UserID:    user.ID,
		User:      user,
		StoreName: u.StoreName,
		City:      u.City,
	}

	return sell
}
