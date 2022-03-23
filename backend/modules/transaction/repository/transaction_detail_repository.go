package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	"gorm.io/gorm"
)

func ReadStaticPrice(tid int, pid int) float64 {
	// Check connection
	con := dbController.GetConnection()

	var staticPrice float64

	// Update object in database
	con.Raw(
		"SELECT `price` "+
			"FROM `transaction_details` "+
			"WHERE transaction_id = ? "+
			"AND product_id = ?",
		tid,
		pid,
	).Scan(&staticPrice)

	return staticPrice
}

// func ReadTDOrderBySeller()  {
// 	// Check connection
// 	con := dbController.GetConnection()

// 	// var

// 	// Update object in database
// 	// result := con.Raw().Scan()

// 	return result
// }

func UpdateTDAfterTransaction(transactionDetails []model.TransactionDetail) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Update object in database
	result := con.Save(&transactionDetails)

	return result
}
