package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
)

func ReadAllAdmins(admId []string) []model.Admin {
	// Check connection
	con := dbController.GetConnection()

	// Get admins from database
	var admins []model.Admin

	// Extra query by id
	if admId != nil {
		con.Find(&admins, admId[0])
	} else {
		con.Find(&admins)
	}

	return admins
}

func CreateAdmin(admin model.Admin) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&admin)

	return result
}

func DeleteAdminByCustId(userId string) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Delete Admin on database
	var adm model.Admin
	result := con.Where("user_id = ?", userId).Delete(&adm)

	return result
}
