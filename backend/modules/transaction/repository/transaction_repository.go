package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	"gorm.io/gorm"
)

func GetTransactionsAssociation(transactions []model.Transaction) {
	// for each transaction
	for i := 0; i < len(transactions); i++ {
		GetOneTransactionAssociation(&transactions[i])
	}
}

func GetOneTransactionAssociation(transaction *model.Transaction) {
	// Check connection
	con := dbController.GetConnection()

	// with product detail
	con.Model(&transaction).Association("ProductDetail").Find(&transaction.ProductDetail)

}

func ReadAllTransaction(transactionId []string) []model.Transaction {
	// Check connection
	con := dbController.GetConnection()

	var transactions []model.Transaction

	// Extra query by id
	if transactionId != nil {
		con.Find(&transactions, transactionId[0])
	} else {
		con.Find(&transactions)
	}

	return transactions
}

func ReadTransactionsByCustId(custId []string) []model.Transaction {
	// Check connection
	con := dbController.GetConnection()

	var transactions []model.Transaction

	// Extra query by id
	if custId != nil {
		con.Where("customer_id = ?", custId[0]).Find(&transactions)
	}

	return transactions
}

func DeleteTransactionById(transactionId string) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	//delete Product on db by id
	var transaction model.Transaction
	result := con.Delete(&transaction, transactionId)
	return result
}

func CreateTransaction(transaction *model.Transaction) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&transaction)

	return result
}
