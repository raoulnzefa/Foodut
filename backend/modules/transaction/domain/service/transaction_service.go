package service

import (
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	repo "github.com/Foodut/backend/modules/transaction/repository"

	// dto "github.com/Foodut/backend/modules/user/rest-api/dto"
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

func DeleteById(transId string) *gorm.DB {
	deleteFeedback := repo.DeleteProductById(transId)

	return deleteFeedback
}
