package service

import (
	"fmt"
	"time"

	prModel "github.com/Foodut/backend/modules/product/domain/model"
	productRepo "github.com/Foodut/backend/modules/product/repository"
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

	sub := GenerateTotalPayment(trans.CustomerId)
	carts := repo.GetCartByCustId(trans.CustomerId)
	var products []prModel.Product

	for i := 0; i < len(carts); i++ {
		products = append(products, productRepo.GetProductById(carts[i].ProductID))
	}

	transaction := model.Transaction{
		CustomerID:      trans.CustomerId,
		PaymentOption:   trans.PaymentOption,
		SubTotal:        sub,
		TransactionDate: GetServerTime(),
		ProductDetail:   products,
	}

	check := repo.CreateTransaction(transaction)
	if check.Error == nil {
		return repo.DeleteCartByCustId(trans.CustomerId)
	}
	return check
}

func DeleteById(transId string) *gorm.DB {
	deleteFeedback := repo.DeleteTransactionById(transId)

	return deleteFeedback
}

func GenerateTotalPayment(customerId int) float64 {
	cart := repo.GetCartByCustId(customerId)
	var subTotal float64
	for i := 0; i < len(cart); i++ {
		product := productRepo.GetProductById(cart[i].ProductID)
		subTotal = subTotal + (float64(cart[i].Quantity) * product.ProductPrice)
	}
	return subTotal
}

func GetServerTime() time.Time {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err)
	}

	return t.In(location)
}
