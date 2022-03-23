package service

import (
	"errors"
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

func InsertTransactionAvailabilityCheck(trans dto.PostTransaction) error {

	// Get customer cart, list only that available in stock
	carts := repo.GetCartByCustIdWithAvailablityCheck(trans.CustomerId)

	// If empty cart,
	// perhaps sold out product,
	// do not checkout
	if len(carts) < 1 {
		err := errors.New("empty cart list. make sure that stock is still available")
		return err
	}

	var products []prModel.Product

	// then read customer cart 🛒from filtered cart list
	for i := 0; i < len(carts); i++ {
		products = append(products, prRepo.ReadProductById(carts[i].ProductID))
	}

	// Generate sub total of cart
	sub := GenerateTotalPaymentDirect(carts, products)

	// Map JSON -> Database Model
	transaction := MapToTransactionModel(trans, sub, products)

	// Insert the transaction to database
	check := repo.CreateTransaction(&transaction)

	// If success
	if check.Error == nil {
		// Update transaction detail quantity
		// Quantity & current price
		if updateTD := SendTDForUpdateAfterTransaction(
			transaction.ID,
			carts,
			products); updateTD.Error != nil {

			// If error, return the error
			return updateTD.Error
		}

		// Update Product Stock due to transaction
		if updatePR := prSrvc.SendForUpdateProductStockAfterTransaction(
			carts,
			products); updatePR.Error == nil {

			// Remove payed product from customer cart
			return repo.DeleteCarts(carts).Error
		}
	}

	return check.Error
}

func DeleteById(transId string) *gorm.DB {
	deleteFeedback := repo.DeleteTransactionById(transId)

	return deleteFeedback
}

func MapToTransactionModel(trans dto.PostTransaction, sub float64, products []prModel.Product) model.Transaction {
	transaction := model.Transaction{
		CustomerID:      trans.CustomerId,
		PaymentOption:   trans.PaymentOption,
		SubTotal:        sub,
		TransactionDate: GetServerTime(),
		ProductDetail:   products,
	}
	return transaction
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

func GenerateTotalPaymentByCustId(customerId int) float64 {
	cart := repo.GetCartByCustId(customerId)
	var subTotal float64
	for i := 0; i < len(cart); i++ {
		product := prRepo.ReadProductById(cart[i].ProductID)
		subTotal = subTotal + (float64(cart[i].Quantity) * product.ProductPrice)
	}
	return subTotal
}

func GenerateTotalPaymentDirect(carts []model.Cart, products []prModel.Product) float64 {
	var subTotal float64
	for i := 0; i < len(carts); i++ {
		subTotal = subTotal + (float64(carts[i].Quantity) * products[i].ProductPrice)
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
