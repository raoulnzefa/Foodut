package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func MapToCustomer(usr dto.PostUser) *gorm.DB {

	// Parse from JSON DTO -> Database Model
	user := MapToUser(usr, 1)

	cust := model.Customer{
		UserID: user.ID,
		User:   user,
	}

	return repo.CreateCustomer(cust)
}
