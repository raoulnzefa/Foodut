package service

import (
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	repo "github.com/Foodut/backend/modules/transaction/repository"
	"github.com/Foodut/backend/modules/transaction/rest-api/dto"

	"gorm.io/gorm"
)

func SearchById(transactionId []string) []model.Transaction {
	transactions := repo.FindAllTransaction(transactionId)

	// Association
	if len(transactions) > 0 {
		repo.GetTransactionsAssociation(transactions)
	}

	return transactions
}

func InsertTransaction(trans dto.Transaction) *gorm.DB {
	transaction := model.Transaction{
		CustomerID:      trans.CustomerId,
		PaymentOption:   trans.PaymentOption,
		TransactionDate: trans.TransactionDate,
	}

	return repo.CreateTransaction(transaction)
}

func DeleteById(transId string) *gorm.DB {
	deleteFeedback := repo.DeleteProductById(transId)

	return deleteFeedback
}
