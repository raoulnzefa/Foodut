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
	deleteCust := repo.DeleteCustomerByCustId(userId)
	if deleteCust.Error == nil {
		deleteAdmn := repo.DeleteAdminByCustId(userId)
		if deleteAdmn.Error == nil {
			deleteSell := repo.DeleteSellerByCustId(userId)
			if deleteSell.Error == nil {
				deleteUser := repo.DeleteUserById(userId)
				return deleteUser
			} else {
				return deleteSell
			}
		} else {
			return deleteAdmn
		}
	} else {
		return deleteCust
	}
}

func CheckUserLogin(email string, password string) (model.User, *gorm.DB) {
	encryptedPassword := GetMD5Hash(password)
	return repo.CheckUserEmailPassword(email, encryptedPassword)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
