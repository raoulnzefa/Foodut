package repository

import (
	dbController "github.com/Foodut/backend/database"
	prRepo "github.com/Foodut/backend/modules/product/repository"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
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

	// get each product association
	prRepo.GetProductsAssociation(transaction.ProductDetail)

}

func FindAllTransaction(transactionId []string) []model.Transaction {
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
