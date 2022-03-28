package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func SendForCreateAdmin(usr dto.PostUser) *gorm.DB {
	return repo.CreateAdmin(MapToAdmin(usr))
}

func MapToAdmin(usr dto.PostUser) model.Admin {

	// Parse from JSON DTO -> Database Model
	user := MapToUser(usr, 3)

	adm := model.Admin{
		UserID: user.ID,
		User:   user,
	}

	return adm
}
