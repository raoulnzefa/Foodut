package service

import (
	"fmt"
	"strconv"
	"time"

	prModel "github.com/Foodut/backend/modules/product/domain/model"
	prSrvc "github.com/Foodut/backend/modules/product/domain/service"
	prRepo "github.com/Foodut/backend/modules/product/repository"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	repo "github.com/Foodut/backend/modules/transaction/repository"
	dto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
	"gorm.io/gorm"
)

func SearchById(transactionId []string) []dto.Transaction {
	transactions := repo.FindAllTransaction(transactionId)

	transId, _ := strconv.ParseInt(transactionId[0], 6, 12)

	var getTransactionDTO []dto.Transaction

	// Association
	if len(transactions) > 0 {
		repo.GetTransactionsAssociation(transactions)

		// Child Association
		for i := 0; i < len(transactions); i++ {
			if len(transactions[i].ProductDetail) > 0 {

				// Map from model to dto
				getTransactionDTO = append(getTransactionDTO, MapToTransactionDTO(transactions[i]))

				// Get from customer
				historyDetail := prSrvc.ProductDetailAssociationWithTransaction(int(transId))

				// Set it to customer collapse
				getTransactionDTO[i].ProductDetail = historyDetail
			}
		}
	}

	return getTransactionDTO
}

func InsertTransaction(trans dto.PostTransaction) *gorm.DB {

	sub := GenerateTotalPayment(trans.CustomerId)
	carts := repo.GetCartByCustId(trans.CustomerId)
	var products []prModel.Product

	for i := 0; i < len(carts); i++ {
		products = append(products, prRepo.GetProductById(carts[i].ProductID))
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

func MapToTransactionDTO(t model.Transaction) dto.Transaction {
	return dto.Transaction{
		ID:              t.ID,
		CustomerId:      t.CustomerID,
		PaymentOption:   t.PaymentOption,
		SubTotal:        t.SubTotal,
		TransactionDate: t.TransactionDate,
	}
}

func GenerateTotalPayment(customerId int) float64 {
	cart := repo.GetCartByCustId(customerId)
	var subTotal float64
	for i := 0; i < len(cart); i++ {
		product := prRepo.GetProductById(cart[i].ProductID)
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
