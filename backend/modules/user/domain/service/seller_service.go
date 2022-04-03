package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func SendForGetSeller() []dto.GetSeller {
	sellers := repo.ReadAllSeller()

	// Association
	if len(sellers) > 0 {
		repo.GetSellersAssociation(sellers)
	}

	var sellerDto []dto.GetSeller

	// Move from database model to JSON DTO
	for _, s := range sellers {
		tempSeller := dto.GetSeller{
			UserID:    s.UserID,
			StoreName: s.StoreName,
			Username:  s.User.Username,
			Email:     s.User.Email,
			Name:      s.User.Name,
			City:      s.City,
		}
		sellerDto = append(sellerDto, tempSeller)
	}

	return sellerDto
}

func SendForGetById(id []string) model.Seller {
	seller := repo.ReadSellerById(id)

	// Association
	if seller.UserID > 0 {
		repo.GetOneSellerAssociation(&seller)
	}

	return seller
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
