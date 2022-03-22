package service

import (
	"crypto/md5"
	"encoding/hex"

	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func SearchUserById(userId []string) []model.User {
	return repo.ReadAllUsers(userId)
}

func MapToUser(usr dto.PostUser, lv int) model.User {
	encryptedPassword := GetMD5Hash(usr.Password)
	// Parse from JSON DTO -> Database Model
	user := model.User{
		Username: usr.Username,
		Email:    usr.Email,
		Name:     usr.Name,
		Password: encryptedPassword,
		Level:    lv,
	}

	return user
}

func DeleteById(userId string) *gorm.DB {
	deleteFeedback := repo.DeleteUserById(userId)

	return deleteFeedback
}

func CheckUserLogin(email string, password string) (int, *gorm.DB) {
	encryptedPassword := GetMD5Hash(password)
	return repo.CheckUserEmailPassword(email, encryptedPassword)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
