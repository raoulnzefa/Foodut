package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func SearchUserById(userId []string) []model.User {
	return repo.ReadAllUsers(userId)
}

func MapToUser(usr dto.PostUser, lv int) model.User {

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

func DeleteById(userId string) *gorm.DB {
	deleteFeedback := repo.DeleteUserById(userId)

	return deleteFeedback
}

func CheckUserLogin(email string, password string) *gorm.DB {
	return repo.CheckUserEmailPassword(email, password)
}
