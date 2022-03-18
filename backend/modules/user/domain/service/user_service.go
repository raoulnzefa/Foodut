package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
)

func EmptyUserSearch() []model.User {
	return repo.FindAllUsers()
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
