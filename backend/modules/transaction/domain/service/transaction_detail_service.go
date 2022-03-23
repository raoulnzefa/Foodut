package service

import (
	prModel "github.com/Foodut/backend/modules/product/domain/model"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	repo "github.com/Foodut/backend/modules/transaction/repository"
	"gorm.io/gorm"
)

func SendTDForUpdateAfterTransaction(tid int, carts []model.Cart, products []prModel.Product) *gorm.DB {
	// Update object in database
	var result *gorm.DB

	var transactionDetails []model.TransactionDetail
	// For each data in carts,
	for i := 0; i < len(carts); i++ {
		transactionDetails = append(transactionDetails, model.TransactionDetail{
			TransactionID: tid,
			ProductID:     products[i].ID,
			Quantity:      carts[i].Quantity,
			Price:         products[i].ProductPrice,
		})
	}

	// record its quantity & price
	result = repo.UpdateTDAfterTransaction(transactionDetails)

	return result
}
