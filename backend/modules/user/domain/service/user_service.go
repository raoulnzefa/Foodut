package service

import (
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
)

func EmptyUserSearch() []model.User {
	return repo.FindAllUsers()
}
