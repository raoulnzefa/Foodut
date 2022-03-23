package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	"gorm.io/gorm"
)

func UpdateTDAfterTransaction(transactionDetails []model.TransactionDetail) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Update object in database
	result := con.Save(&transactionDetails)

	return result
}
