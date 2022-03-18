package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func EmptyUserSearch() []model.User {
	return repo.FindAllUsers()
}

func mapToUser(usr dto.PostUser, lv int) model.User {

	// Parse from JSON DTO -> Database Model
	user := model.User{
		Username: usr.Username,
		Email:    usr.Email,
		Name:     usr.Name,
		Password: usr.Password,
		Level:    lv,
	}

	return user
}

func MapToCustomer(usr dto.PostUser) *gorm.DB {

	// Parse from JSON DTO -> Database Model
	user := mapToUser(usr, 1)

	cust := model.Customer{
		UserID: user.ID,
		User:   user,
	}

	return repo.CreateCustomer(cust)
}

func MapToSeller(usr dto.PostUser) *gorm.DB {

	// Parse from JSON DTO -> Database Model
	user := mapToUser(usr, 2)

	sell := model.Seller{
		UserID: user.ID,
		User:   user,
	}

	return repo.CreateSeller(sell)
}

func MapToAdmin(usr dto.PostUser) *gorm.DB {

	// Parse from JSON DTO -> Database Model
	user := mapToUser(usr, 3)

	adm := model.Admin{
		UserID: user.ID,
		User:   user,
	}

	return repo.CreateAdmin(adm)
}
