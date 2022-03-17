package repository

import (
	"fmt"

	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	"gorm.io/gorm"
)

func FindAllTransaction(transactionId []string) []model.Transaction {
	// Check connection
	con := dbController.GetConnection()
	var result *gorm.DB

	var transactions []model.Transaction

	// Extra query by id
	if transactionId != nil {
		result = con.Find(&transactions, transactionId[0])
	} else {
		result = con.Find(&transactions)
	}

	// Association
	if result.Error == nil && len(transactions) > 0 {
		// for each transaction
		for i := 0; i < len(transactions); i++ {
			// with product detail
			con.Model(&transactions).Association("ProductDetail").Find(&transactions[i].ProductDetail)

			// with picture
			for j := 0; j < len(transactions[i].ProductDetail); j++ {
				con.Model(&transactions[i].ProductDetail).Association("Picture").Find(&transactions[i].ProductDetail[j].Picture)
			}
		}
	} else {
		fmt.Println(result.Error)
	}

	return transactions
}
