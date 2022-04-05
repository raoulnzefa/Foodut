package service

import (
	prModel "github.com/Foodut/backend/modules/product/domain/model"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	repo "github.com/Foodut/backend/modules/transaction/repository"
	"github.com/Foodut/backend/modules/transaction/rest-api/dto"
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
			SellerID:      products[i].SellerID,
			Quantity:      carts[i].Quantity,
			Price:         products[i].ProductPrice,
			Status:        "ORDER",
		})
	}

	// record its quantity & price
	result = repo.UpdateTDAfterTransaction(transactionDetails)

	return result
}

func SendForGetTransactionDetail() []dto.TransactionDecentralize {
	td := repo.ReadAllTransactionDetail()

	// Calculate each product sub total
	for i := 0; i < len(td); i++ {
		td[i].SubTotal = td[i].Price * float64(td[i].Quantity)
	}

	return td
}

func SendSellerIdForReadOrder(sellerId []string) []dto.OrderDetail {
	return repo.ReadTDOrderBySeller(sellerId)
}

func SendCustomerIdForReadOrder(customerId []string) []dto.OrderDetail {
	return repo.ReadTDOrderBySeller(customerId)
}
