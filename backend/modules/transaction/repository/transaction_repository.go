package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
)

func FindAllTransaction(transactionId []string) []model.Transaction {
	// Check connection
	con := dbController.GetConnection()

	var transactions []model.Transaction

	if transactionId != nil {
		con.Find(&transactions, transactionId[0])
	} else {
		con.Find(&transactions)
	}

	for i := 0; i < len(transactions); i++ {
		con.Model(&transactions[i]).Association("ProductDetail").Find(&transactions[i].ProductDetail)

		con.Model(&transactions[i].ProductDetail).Association("Picture")
		for j := 0; j < len(transactions[i].ProductDetail); j++ {
			con.Model(&transactions[i].ProductDetail[j]).Association("Picture").Find(&transactions[i].ProductDetail[j].Picture)
		}
	}

	return transactions
}
