package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func MapToAdmin(usr dto.PostUser) *gorm.DB {

	// Parse from JSON DTO -> Database Model
	user := MapToUser(usr, 3)

	adm := model.Admin{
		UserID: user.ID,
		User:   user,
	}

	return repo.CreateAdmin(adm)
}
