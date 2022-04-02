package service

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

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
			deleteSell := repo.DeleteSellerBySellerId(userId)
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

func EditUser(user model.User, editUserDto dto.EditUser) (*gorm.DB, *gorm.DB, int) {
	levelUser := user.Level // buat dipake saat cek level user apakah cust atau seller
	username := editUserDto.Username
	name := editUserDto.Name
	password := editUserDto.Password
	profilePhoto := editUserDto.ProfilePhoto
	addressCust := editUserDto.Address
	citySeller := editUserDto.City
	storeNameSeller := editUserDto.StoreName

	if username != "" {
		user.Username = username
	}
	if name != "" {
		user.Name = name
	}
	if password != "" {
		encryptedPassword := GetMD5Hash(password)
		user.Password = encryptedPassword
	}
	if profilePhoto != "" {
		user.ProfilePhoto = profilePhoto
	}
	if user.Level == 1 {
		var id []string
		id = append(id, strconv.Itoa(user.ID))
		custs := repo.ReadAllCustomers(id)
		if addressCust != "" {
			custs[0].Address = addressCust
		}

		if len(custs) > 0 {
			result1 := repo.SaveUser(user)
			result2 := repo.SaveCustomer(custs[0])
			return result1, result2, 1
		} else {
			return nil, nil, 1
		}
	} else if levelUser == 2 {
		var id []string
		id = append(id, strconv.Itoa(user.ID))
		sellers := repo.ReadAllSellers(id)
		if citySeller != "" {
			sellers[0].City = citySeller
		}
		if storeNameSeller != "" {
			sellers[0].StoreName = storeNameSeller
		}
		if len(sellers) > 0 {
			result1 := repo.SaveUser(user)
			result2 := repo.SaveSeller(sellers[0])
			return result1, result2, 2
		} else {
			return nil, nil, 2
		}
	}
	return nil, nil, 3
}
